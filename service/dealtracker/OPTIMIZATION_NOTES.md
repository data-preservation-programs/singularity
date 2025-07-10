# Deal Tracker Performance Optimizations

## Overview

This document describes the performance optimizations implemented for the deal tracker to improve throughput from 1.8MB/s to potentially 20+ MB/s.

## Problem Analysis

From the production metrics:
- 31 million voluntary context switches in 20 minutes (~26,000/second)
- Only using ~2 CPU cores (194% CPU usage)
- Very low memory usage (110MB)
- Processing bottlenecked by excessive goroutine scheduling and DB contention

## Optimizations Implemented

### 1. Fast-Path Client Filtering

**Before**: Every deal was fully decoded using `mapstructure.Decode()` before checking if it belonged to our wallets.

**After**: Two-phase approach:
- Phase 1: Quick check of `Proposal.Client` field without full parsing
- Phase 2: Full parsing only for matching deals

**Implementation**:
- Type assertion fast path for `map[string]any` values (9.3ns per check)
- `gjson` fallback for other formats (161ns per check)
- Skip ~90%+ of deals without expensive parsing

### 2. Batch Processing

**Before**: Each deal processed individually, causing high DB contention.

**After**: Process deals in configurable batches (default 100).

**Benefits**:
- Reduced context switches
- Better CPU cache utilization
- Opportunity for batch DB operations

### 3. Streaming Architecture

Maintained low memory usage by:
- Processing deals as they stream in
- Never holding more than one batch in memory
- Configurable batch size for memory/performance tradeoff

## Performance Characteristics

### Benchmarks

```
BenchmarkBatchParser_ShouldProcessDeal/MapInterface    9.3ns/op     0 B/op    0 allocs/op
BenchmarkBatchParser_ShouldProcessDeal/JSONBytes       161ns/op    32 B/op    2 allocs/op
BenchmarkBatchParser_ParseDeal                        9.4μs/op  8715 B/op  125 allocs/op
```

### Expected Improvements

- **Filtering**: ~1000x faster for non-matching deals
- **Throughput**: 3-4x improvement from fast-path alone
- **CPU Utilization**: Better with batching and reduced contention
- **Memory**: Still under 500MB constraint

## Configuration

The batch size can be configured when creating the DealTracker. Default is 100 deals per batch.

```go
tracker := NewDealTracker(db, interval, dealURL, lotusURL, lotusToken, once)
tracker.batchSize = 200 // Adjust based on your needs
```

## Future Optimizations

1. **Parallel Processing**: Process multiple batches concurrently
2. **Custom JSON Parser**: Skip `map[string]interface{}` entirely
3. **Batch DB Operations**: Batch inserts/updates to reduce DB round trips
4. **Connection Pooling**: Better DB connection management 