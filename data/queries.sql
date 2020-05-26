--CREATE DATABASE 
    CREATE DATABASE IF NOT EXISTS driverdemo;

--CREATE TABLE
    CREATE EXTERNAL TABLE `peopletable`(
    `first_name` string, 
    `last_name` string, 
    `company_name` string, 
    `address` string, 
    `city` string, 
    `state` string, 
    `post` bigint, 
    `phone1` string, 
    `phone2` string, 
    `email` string, 
    `web` string)
    ROW FORMAT DELIMITED 
    FIELDS TERMINATED BY ',' 
    STORED AS INPUTFORMAT 
    'org.apache.hadoop.mapred.TextInputFormat' 
    OUTPUTFORMAT 
    'org.apache.hadoop.hive.ql.io.HiveIgnoreKeyTextOutputFormat'
    LOCATION
    's3://brads-playground/Athena/DemoCSV/'

--SELECT FROM TABLE
    SELECT *
    FROM "peopledb"."peopletable"

--SELECT QUERY (NO WHERE CLAUSE)
    SELECT State,
    Count(*) as "Count"

    FROM "peopledb"."peopletable"

    GROUP BY State

--SELECT QUERY (WHERE CLAUSE)
    SELECT State,
    Count(*) as "Count"

    FROM "peopledb"."peopletable"

    WHERE State = 'VIC'

    GROUP BY State


