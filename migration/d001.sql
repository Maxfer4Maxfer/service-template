START TRANSACTION;

SET ROLE advertisement;

--------------------------------------------------------------------------
DROP TABLE IF EXISTS technical.version;
DROP SCHEMA IF EXISTS technical;
--------------------------------------------------------------------------

DROP TABLE IF EXISTS advert_history;
DROP TABLE IF EXISTS advert;
DROP TABLE IF EXISTS advert_status;
DROP TABLE IF EXISTS advert_type;
DROP TABLE IF EXISTS campaign_history;
DROP TABLE IF EXISTS campaign;
DROP TABLE IF EXISTS campaign_status;
DROP TABLE IF EXISTS supplier;

COMMIT TRANSACTION;
-- ROLLBACK TRANSACTION;

--------------------------------------------------------------------------
SET ROLE master;
DROP USER IF EXISTS advertisement;
--------------------------------------------------------------------------
