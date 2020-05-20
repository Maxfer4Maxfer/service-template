CREATE DATABASE calculator;
CREATE USER calculator WITH ENCRYPTED PASSWORD 'calculator';
GRANT ALL PRIVILEGES ON DATABASE calculator TO calculator;
--------------------------------------------------------------------------

START TRANSACTION;
SET ROLE calculator;

--------------------------------------------------------------------------
-- log -------------------------------------------------------------------
--------------------------------------------------------------------------
CREATE TABLE public.log (
  id serial NOT NULL PRIMARY KEY,
  created_at timestamptz NOT NULL,
  operation_type varchar(255),
  log_text varchar(255)
);

COMMENT ON TABLE public.log is 'Log of calculation operations';
COMMENT ON COLUMN public.log.created_at IS 'Timestamp';
COMMENT ON COLUMN public.log.operation_type IS 'Type of calc operation';
COMMENT ON COLUMN public.log.log_text IS 'Text of a log record';

ALTER TABLE public.log OWNER TO calculator;

--------------------------------------------------------------------------
-- log_warehouse ---------------------------------------------------------
--------------------------------------------------------------------------
CREATE TABLE public.log_warehouse (
  id serial NOT NULL PRIMARY KEY,
  created_at timestamptz NOT NULL,
  operation_type varchar(255),
  log_text varchar(255)
);

COMMENT ON TABLE public.log_warehouse is 'Log of calculation operations';
COMMENT ON COLUMN public.log_warehouse.created_at IS 'Timestamp';
COMMENT ON COLUMN public.log_warehouse.operation_type IS 'Type of calc operation';
COMMENT ON COLUMN public.log_warehouse.log_text IS 'Text of a log record';

ALTER TABLE public.log_warehouse OWNER TO calculator;
--------------------------------------------------------------------------
--------------------------------------------------------------------------

------------------------------technical----------------------------------
CREATE SCHEMA technical;
CREATE TABLE technical.version
(
    version_num int NOT NULL,
    apply_date timestamptz NULL
);

COMMENT ON TABLE technical.version is 'Schema version';

INSERT INTO technical.version
    ( version_num, apply_date )
VALUES
    ( 1, now() );
--------------------------------------------------------------------------


 COMMIT TRANSACTION;
 --ROLLBACK TRANSACTION;
