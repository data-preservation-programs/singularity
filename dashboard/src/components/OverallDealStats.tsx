import React, {useEffect, useState} from 'react';
import axios from 'axios';
import {Card, CardGroup} from 'react-bootstrap';
import {
    Bar,
    BarChart,
    CartesianGrid,
    Legend,
    Line,
    LineChart,
    ResponsiveContainer,
    Tooltip,
    XAxis,
    YAxis,
} from 'recharts';
import xbytes from 'xbytes';

interface DealStats {
  Provider: string;
  State: string;
  Day: string;
  DealSize: number;
}

const CustomYAxisTickFormatter = (value: number) => {
  return xbytes(value, {iec: true});
};
const CustomXAxisTickFormatter = (value: any, index: number) => {
  const date = new Date(value);
  const year = date.getFullYear();
  const month = ("0" + (date.getMonth() + 1)).slice(-2);
  const day = ("0" + date.getDate()).slice(-2);
  return `${year}-${month}-${day}`;
};

const OverallDealStats: React.FC = () => {
  const [dealStats, setDealStats] = useState<DealStats[]>([]);

  useEffect(() => {
    axios.get(`${process.env.REACT_APP_API_BASE_URL}/api/deal_stats`).then((response) => {
      setDealStats(response.data);
    });
  }, []);

  // Calculate required stats from the dealStats
  const activeDeals = dealStats.filter((deal) => deal.State === 'active');
  const totalActiveDealsNumber = activeDeals.reduce((acc, deal) => acc + deal.DealSize, 0)
  const totalActiveDeals = xbytes(totalActiveDealsNumber, {iec: true});
  const qap = xbytes(totalActiveDealsNumber * 10, {iec: true})
  const distinctProviders = new Set(activeDeals.map((deal) => deal.Provider)).size;

  // Prepare data for the time series line chart
  const timeSeriesDataMap = activeDeals.reduce((acc, deal) => {
    const day = deal.Day;
    if (!acc.has(day)) {
      acc.set(day, { Day: day, ActiveDeals: 0 });
    }
    acc.get(day).ActiveDeals+= deal.DealSize;
    return acc;
  }, new Map());

  const timeSeriesData = Array.from(timeSeriesDataMap.values()).sort((a, b) => a.Day.localeCompare(b.Day));

// Calculate the cumulative active deals count
  let cumulativeCount = 0;
  for (let i = 0; i < timeSeriesData.length; i++) {
    cumulativeCount += timeSeriesData[i].ActiveDeals;
    timeSeriesData[i].ActiveDeals = cumulativeCount;
  }

  // Prepare data for the bar charts
  const dealsByStateDataMap = new Map<string, number>();
  for (const deal of dealStats) {
    const state = deal.State;
    if (!dealsByStateDataMap.has(state)) {
      dealsByStateDataMap.set(state, 0);
    }
    dealsByStateDataMap.set(state, dealsByStateDataMap.get(state)! + deal.DealSize);
  }
  const dealsByStateData = Array.from(dealsByStateDataMap.entries())
    .map(([State, Deals]) => ({ State, Deals }))
    .sort((a, b) => b.Deals - a.Deals);

  const activeDealsByProviderDataMap = new Map<string, number>()
  for (const deal of activeDeals) {
    const provider = deal.Provider;
    if (!activeDealsByProviderDataMap.has(provider)) {
      activeDealsByProviderDataMap.set(provider, 0);
    }
    activeDealsByProviderDataMap.set(provider, activeDealsByProviderDataMap.get(provider)! + deal.DealSize);
  }

  const activeDealsByProviderData = Array.from(activeDealsByProviderDataMap.entries())
    .map(([Provider, Deals]) => ({ Provider, Deals }))
    .sort((a, b) => b.Deals - a.Deals);


  return (
    <div>
      {/* First row - Total active deals and distinct providers */}
      <CardGroup>
        <Card>
          <Card.Body>
            <Card.Title>Total Active Deals</Card.Title>
            <Card.Text>{totalActiveDeals}</Card.Text>
          </Card.Body>
        </Card>
        <Card>
          <Card.Body>
            <Card.Title>Quality Adjusted Power</Card.Title>
            <Card.Text>{qap}</Card.Text>
          </Card.Body>
        </Card>
        <Card>
          <Card.Body>
            <Card.Title>Distinct Providers</Card.Title>
            <Card.Text>{distinctProviders}</Card.Text>
          </Card.Body>
        </Card>
      </CardGroup>

  {/* Second row - Time series line chart */}
  <div style={{ marginTop: '1rem', marginBottom: '1rem', width: '100%', height: '400px' }}>
    <ResponsiveContainer>
      <LineChart data={timeSeriesData} margin={{top: 20, left: 20}}>
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="Day" tickFormatter={CustomXAxisTickFormatter} />
        <YAxis tickFormatter={CustomYAxisTickFormatter} />
        <Tooltip />
        <Legend />
        <Line type="monotone" dataKey="ActiveDeals" stroke="#8884d8" />
      </LineChart>
    </ResponsiveContainer>
  </div>
      {/* Third row - Bar charts */}
      <div style={{ display: 'flex', justifyContent: 'space-between' }}>
        <div style={{ width: '45%', height: `${activeDealsByProviderData.length * 30}px` }}>
          <ResponsiveContainer>
            <BarChart layout="vertical" data={activeDealsByProviderData} margin={{left: 25}}>
              <CartesianGrid strokeDasharray="3 3" />
              <YAxis type="category" dataKey="Provider" />
              <XAxis type="number" tickFormatter={CustomYAxisTickFormatter} orientation="top" />
              <Tooltip />
              <Legend verticalAlign="top"  />
              <Bar dataKey="Deals" name="Deals by state" fill="#8884d8" />
            </BarChart>
          </ResponsiveContainer>
        </div>
        <div style={{ width: '45%', height: `${dealsByStateData.length * 50}px` }}>
          <ResponsiveContainer>
            <BarChart layout="vertical" data={dealsByStateData} margin={{left: 70}} >
              <CartesianGrid strokeDasharray="3 3" />
              <YAxis type="category" dataKey="State" />
              <XAxis type="number" tickFormatter={CustomYAxisTickFormatter} orientation="top" />
              <Tooltip />
              <Legend verticalAlign="top" />
              <Bar dataKey="Deals" name="Active deals by Provider" fill="#82ca9d" />
            </BarChart>
          </ResponsiveContainer>
        </div>
      </div>
    </div>
  );
};

export default OverallDealStats;
