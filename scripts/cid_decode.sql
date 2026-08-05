-- Decode singularity's bytea CID columns (the raw cid.Bytes() form, e.g.
-- car_blocks.cid, cars.piece_cid, cars.root_cid) into the canonical CID
-- string, matching go-cid's Cid.String(). Pure PL/pgSQL, no extension needed.
--   CIDv1 -> multibase base32 with a leading 'b'  (bafy…/bafk…/baga…)
--   CIDv0 -> base58btc, no prefix                 (Qm…)
--
-- usage:
--   SELECT cid_str(cid) FROM car_blocks WHERE file_id = 123 LIMIT 20;
--   SELECT cid_str(piece_cid), cid_str(root_cid) FROM cars WHERE id = 456;
-- cid_str is IMMUTABLE, so it can also back a functional index or a
-- GENERATED ALWAYS AS (cid_str(cid)) STORED column. It's per-row PL/pgSQL,
-- so filter (WHERE/LIMIT) rather than scanning all of car_blocks.

-- base32, lowercase, RFC 4648, no padding (CIDv1 multibase 'b')
CREATE OR REPLACE FUNCTION cid_b32(data bytea) RETURNS text
LANGUAGE plpgsql IMMUTABLE STRICT AS $$
DECLARE
  alpha CONSTANT text := 'abcdefghijklmnopqrstuvwxyz234567';
  out text := '';
  buf  int  := 0;
  bits int  := 0;
  i    int;
BEGIN
  FOR i IN 0 .. length(data)-1 LOOP
    buf  := (buf << 8) | get_byte(data, i);
    bits := bits + 8;
    WHILE bits >= 5 LOOP
      bits := bits - 5;
      out  := out || substr(alpha, ((buf >> bits) & 31) + 1, 1);
    END LOOP;
    buf := buf & ((1 << bits) - 1);   -- keep only the leftover bits (avoid int overflow)
  END LOOP;
  IF bits > 0 THEN
    out := out || substr(alpha, ((buf << (5 - bits)) & 31) + 1, 1);
  END IF;
  RETURN out;
END $$;

-- base58btc (Bitcoin alphabet)
CREATE OR REPLACE FUNCTION cid_b58(data bytea) RETURNS text
LANGUAGE plpgsql IMMUTABLE STRICT AS $$
DECLARE
  alpha CONSTANT text := '123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz';
  n     numeric := 0;
  out   text := '';
  i     int;
  zeros int := 0;
BEGIN
  WHILE zeros < length(data) AND get_byte(data, zeros) = 0 LOOP   -- leading 0x00 -> '1'
    zeros := zeros + 1;
  END LOOP;
  FOR i IN 0 .. length(data)-1 LOOP                               -- big-endian bytes -> bignum
    n := n * 256 + get_byte(data, i);
  END LOOP;
  WHILE n > 0 LOOP
    out := substr(alpha, (n % 58)::int + 1, 1) || out;
    n   := div(n, 58);
  END LOOP;
  RETURN repeat('1', zeros) || out;
END $$;

-- main entry point: bytea CID -> canonical string
CREATE OR REPLACE FUNCTION cid_str(data bytea) RETURNS text
LANGUAGE plpgsql IMMUTABLE STRICT AS $$
BEGIN
  IF length(data) = 0 THEN
    RETURN NULL;
  END IF;
  -- CIDv0: bare 34-byte sha2-256 multihash (0x12 0x20 …)
  IF length(data) = 34 AND get_byte(data,0) = 18 AND get_byte(data,1) = 32 THEN
    RETURN cid_b58(data);
  END IF;
  -- CIDv1: version byte 0x01 -> base32 multibase
  IF get_byte(data,0) = 1 THEN
    RETURN 'b' || cid_b32(data);
  END IF;
  RETURN NULL;   -- not a CID we recognize
END $$;
