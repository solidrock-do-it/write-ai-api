-- Create material_providers table (match GORM default pluralized name)
CREATE TABLE IF NOT EXISTS material_providers (
    name VARCHAR(255) PRIMARY KEY,
    keys TEXT
);