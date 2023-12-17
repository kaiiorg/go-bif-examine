SELECT 
    'CREATE DATABASE go_bif_examine'
WHERE 
    NOT EXISTS (SELECT FROM pg_database WHERE datname = 'go_bif_examine');
