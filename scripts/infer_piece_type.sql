-- Infer piece_type for cars that predate the piece_type column
-- A piece is Data if any of its blocks reference files (contain file content)
-- A piece is DAG if none of its blocks reference files (contain only directory metadata)

-- Preview what would be updated
SELECT
    c.id,
    c.piece_cid,
    c.piece_size,
    c.piece_type as current_type,
    CASE
        WHEN EXISTS (
            SELECT 1 FROM car_blocks cb
            WHERE cb.car_id = c.id AND cb.file_id IS NOT NULL
        ) THEN 'data'
        ELSE 'dag'
    END as inferred_type
FROM cars c
WHERE c.piece_type IS NULL OR c.piece_type = '';

-- Uncomment to actually update:
-- UPDATE cars c
-- SET piece_type = CASE
--     WHEN EXISTS (
--         SELECT 1 FROM car_blocks cb
--         WHERE cb.car_id = c.id AND cb.file_id IS NOT NULL
--     ) THEN 'data'
--     ELSE 'dag'
-- END
-- WHERE c.piece_type IS NULL OR c.piece_type = '';
