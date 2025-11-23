-- ============================
-- ENUM DEFINITIONS
-- ============================

-- user roles
CREATE TYPE user_role AS ENUM ('admin', 'accountant');

-- site status
CREATE TYPE site_status AS ENUM ('active', 'completed');

-- payout/bill status
CREATE TYPE record_status AS ENUM ('draft', 'saved', 'paid', 'partial');

-- attendance status
CREATE TYPE attendance_status AS ENUM ('present', 'absent', 'half_day');

-- ============================
-- TABLES
-- ============================

-- 1. users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    role user_role NOT NULL DEFAULT 'admin',
    password_hash VARCHAR(255) NOT NULL,
    additional_details JSONB DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- 2. workers
CREATE TABLE workers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(50) UNIQUE,
    skill VARCHAR(100),
    bill_rate NUMERIC(10,2) NOT NULL CHECK (bill_rate >= 0),
    payout_rate NUMERIC(10,2) NOT NULL CHECK (payout_rate >= 0),
    active_status BOOLEAN NOT NULL DEFAULT TRUE,
    notes TEXT,
    additional_details JSONB DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- 3. sites
CREATE TABLE sites (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    client_name VARCHAR(255) NOT NULL,
    client_phone VARCHAR(50),
    location VARCHAR(255),
    site_documents TEXT,
    additional_details JSONB DEFAULT '{}'::jsonb,
    start_date DATE NOT NULL,
    end_date DATE,
    status site_status NOT NULL DEFAULT 'active',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- 4. attendance
CREATE TABLE attendance (
    id SERIAL PRIMARY KEY,
    worker_id INT NOT NULL REFERENCES workers(id) ON DELETE CASCADE,
    site_id INT NOT NULL REFERENCES sites(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    status attendance_status NOT NULL,
    note TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    -- worker cannot have 2 attendance entries for same day in same site
    UNIQUE(worker_id, site_id, date)
);

-- 5. advances
CREATE TABLE advances (
    id SERIAL PRIMARY KEY,
    worker_id INT NOT NULL REFERENCES workers(id) ON DELETE CASCADE,
    amount NUMERIC(10,2) NOT NULL CHECK (amount > 0),
    given_on DATE NOT NULL,
    note TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- 6. payouts
CREATE TABLE payouts (
    id SERIAL PRIMARY KEY,
    period_start DATE NOT NULL,
    period_end DATE NOT NULL,
    payout_json JSONB NOT NULL,
    status record_status NOT NULL DEFAULT 'saved',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- 7. advance_applications
CREATE TABLE advance_applications (
    id SERIAL PRIMARY KEY,
    worker_id INT NOT NULL REFERENCES workers(id) ON DELETE CASCADE,
    payout_id INT NOT NULL REFERENCES payouts(id) ON DELETE CASCADE,
    applied_amount NUMERIC(10,2) NOT NULL CHECK (applied_amount >= 0),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- 8. bills
CREATE TABLE bills (
    id SERIAL PRIMARY KEY,
    site_id INT NOT NULL REFERENCES sites(id) ON DELETE CASCADE,
    period_start DATE NOT NULL,
    period_end DATE NOT NULL,
    bill_json JSONB NOT NULL,
    status record_status NOT NULL DEFAULT 'saved',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- 9. site_summary_snapshots
CREATE TABLE site_summary_snapshots (
    id SERIAL PRIMARY KEY,
    site_id INT NOT NULL REFERENCES sites(id) ON DELETE CASCADE,
    summary_json JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
