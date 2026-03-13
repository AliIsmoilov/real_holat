-- 1. ENUMS
CREATE TYPE user_role AS ENUM ('citizen', 'moderator', 'admin');
CREATE TYPE infra_status AS ENUM ('green', 'yellow', 'red');

-- 2. AUTOMATIC UPDATED_AT FUNCTION
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- 3. INFRASTRUCTURE TYPES
CREATE TABLE infrastructure_types (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    icon_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER update_infra_types_updated_at BEFORE UPDATE ON infrastructure_types FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

-- 4. USERS TABLE
CREATE TABLE users (
    id UUID PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) UNIQUE NOT NULL,
    role user_role DEFAULT 'citizen',
    tg_id BIGINT,
    tg_user_name VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

-- 5. INFRASTRUCTURES TABLE
CREATE TABLE infrastructures (
    id UUID PRIMARY KEY,
    type_id UUID REFERENCES infrastructure_types(id),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    address TEXT NOT NULL,
    latitude DECIMAL(9,6) NOT NULL,
    longitude DECIMAL(9,6) NOT NULL,
    current_status infra_status DEFAULT 'green',
    overall_rating INTEGER DEFAULT 100,
    contractor_name VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER update_infrastructures_updated_at BEFORE UPDATE ON infrastructures FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

-- 6. INFRASTRUCTURE CHECK ITEMS
CREATE TABLE infrastructure_check_items (
    id UUID PRIMARY KEY,
    infrastructure_id UUID NOT NULL REFERENCES infrastructures(id) ON DELETE CASCADE,
    category VARCHAR(50), 
    question TEXT NOT NULL,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER update_check_items_updated_at BEFORE UPDATE ON infrastructure_check_items FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

-- 7. REPORTS TABLE
CREATE TABLE reports (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    infrastructure_id UUID NOT NULL REFERENCES infrastructures(id) ON DELETE CASCADE,
    photo_url TEXT NOT NULL, 
    comment TEXT,
    lat_at_submission DECIMAL(9,6),
    long_at_submission DECIMAL(9,6),
    is_verified BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER update_reports_updated_at BEFORE UPDATE ON reports FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

-- 8. REPORT DETAILS
CREATE TABLE report_details (
    id UUID PRIMARY KEY,
    report_id UUID NOT NULL REFERENCES reports(id) ON DELETE CASCADE,
    check_item_id UUID NOT NULL REFERENCES infrastructure_check_items(id) ON DELETE CASCADE,
    answer BOOLEAN NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    UNIQUE(report_id, check_item_id)
);
CREATE TRIGGER update_report_details_updated_at BEFORE UPDATE ON report_details FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

-- 9. INDEXES (Optimized for Soft Deletes)
CREATE INDEX idx_infra_active ON infrastructures(latitude, longitude) WHERE deleted_at IS NULL;
CREATE INDEX idx_reports_active ON reports(infrastructure_id) WHERE deleted_at IS NULL;