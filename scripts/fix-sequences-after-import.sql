-- reset sequences after importing data with explicit IDs (e.g., from mysql)
-- postgresql sequences don't auto-update when inserting with explicit IDs
-- setval(seq, max, true) means next nextval() returns max+1

SELECT setval(pg_get_serial_sequence('preparations', 'id'), COALESCE((SELECT MAX(id) FROM preparations), 0), true);
SELECT setval(pg_get_serial_sequence('storages', 'id'), COALESCE((SELECT MAX(id) FROM storages), 0), true);
SELECT setval(pg_get_serial_sequence('source_attachments', 'id'), COALESCE((SELECT MAX(id) FROM source_attachments), 0), true);
SELECT setval(pg_get_serial_sequence('output_attachments', 'id'), COALESCE((SELECT MAX(id) FROM output_attachments), 0), true);
SELECT setval(pg_get_serial_sequence('jobs', 'id'), COALESCE((SELECT MAX(id) FROM jobs), 0), true);
SELECT setval(pg_get_serial_sequence('files', 'id'), COALESCE((SELECT MAX(id) FROM files), 0), true);
SELECT setval(pg_get_serial_sequence('file_ranges', 'id'), COALESCE((SELECT MAX(id) FROM file_ranges), 0), true);
SELECT setval(pg_get_serial_sequence('directories', 'id'), COALESCE((SELECT MAX(id) FROM directories), 0), true);
SELECT setval(pg_get_serial_sequence('cars', 'id'), COALESCE((SELECT MAX(id) FROM cars), 0), true);
SELECT setval(pg_get_serial_sequence('car_blocks', 'id'), COALESCE((SELECT MAX(id) FROM car_blocks), 0), true);
SELECT setval(pg_get_serial_sequence('deals', 'id'), COALESCE((SELECT MAX(id) FROM deals), 0), true);
SELECT setval(pg_get_serial_sequence('schedules', 'id'), COALESCE((SELECT MAX(id) FROM schedules), 0), true);
-- workers, globals, wallets have string PKs (no sequence)
