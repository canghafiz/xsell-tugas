-- Drop existing types if they exist
DROP TYPE IF EXISTS Role CASCADE;
DROP TYPE IF EXISTS ProductCondition CASCADE;
DROP TYPE IF EXISTS ProductStatus CASCADE;
DROP TYPE IF EXISTS BookingStatus CASCADE;
DROP TYPE IF EXISTS PaymentStatus CASCADE;
DROP TYPE IF EXISTS BalanceType CASCADE;

-- Create ENUM types
CREATE TYPE Role AS ENUM ('Member', 'Admin');
CREATE TYPE ProductCondition AS ENUM ('New', 'Like New', 'Good', 'Good Quite', 'Needs Repair');
CREATE TYPE ProductStatus AS ENUM ('Available', 'Sold out');
CREATE TYPE BookingStatus AS ENUM ('Pending', 'Waiting confirmation', 'Accept', 'Decline');
CREATE TYPE PaymentStatus AS ENUM ('Pending', 'Success', 'Failed', 'Expire');
CREATE TYPE BalanceType AS ENUM ('Top Up', 'Withdraw');

-- Drop existing tables (in reverse dependency order)
DROP TABLE IF EXISTS Messages CASCADE;
DROP TABLE IF EXISTS Banners CASCADE;
DROP TABLE IF EXISTS UserBalanceHistory CASCADE;
DROP TABLE IF EXISTS UserBalances CASCADE;
DROP TABLE IF EXISTS Payments CASCADE;
DROP TABLE IF EXISTS Bookings CASCADE;
DROP TABLE IF EXISTS Wishlists CASCADE;
DROP TABLE IF EXISTS ProductImages CASCADE;
DROP TABLE IF EXISTS ProductSpecs CASCADE;
DROP TABLE IF EXISTS CategoryProductSpecs CASCADE;
DROP TABLE IF EXISTS Products CASCADE;
DROP TABLE IF EXISTS Location CASCADE;
DROP TABLE IF EXISTS SubCategories CASCADE;  -- new
DROP TABLE IF EXISTS Categories CASCADE;
DROP TABLE IF EXISTS UserVerified CASCADE;
DROP TABLE IF EXISTS Users CASCADE;

-- Create Users table
CREATE TABLE Users (
                       user_id SERIAL PRIMARY KEY,
                       role Role DEFAULT 'Member',
                       first_name VARCHAR(25) NOT NULL,
                       last_name VARCHAR(25),
                       email VARCHAR(100) NOT NULL UNIQUE,
                       password VARCHAR(100) NOT NULL,
                       photo_profile_url TEXT NULL,
                       token TEXT NULL,
                       token_expire TIMESTAMP NULL DEFAULT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create UserVerified table
CREATE TABLE UserVerified (
                              verified_id SERIAL PRIMARY KEY,
                              user_id INTEGER REFERENCES Users(user_id) ON DELETE CASCADE,
                              email VARCHAR(50),
                              verified BOOLEAN DEFAULT FALSE,
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create Categories table (main categories only)
CREATE TABLE Categories (
                            category_id SERIAL PRIMARY KEY,
                            category_name VARCHAR(50) NOT NULL UNIQUE,
                            category_slug VARCHAR(50) NOT NULL UNIQUE,
                            description TEXT,
                            icon VARCHAR(255),
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create SubCategories table
CREATE TABLE SubCategories (
                               sub_category_id SERIAL PRIMARY KEY,
                               parent_category_id INTEGER NOT NULL REFERENCES Categories(category_id) ON DELETE CASCADE,
                               sub_category_name VARCHAR(50) NOT NULL,
                               sub_category_slug VARCHAR(50) NOT NULL UNIQUE,
                               description TEXT,
                               icon VARCHAR(255),
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               UNIQUE (parent_category_id, sub_category_name)
);

-- Create Products table (now linked to SubCategories)
CREATE TABLE Products (
                          product_id SERIAL PRIMARY KEY,
                          listing_user_id INTEGER REFERENCES Users(user_id) ON DELETE SET NULL,
                          sub_category_id INTEGER NOT NULL REFERENCES SubCategories(sub_category_id) ON DELETE RESTRICT,
                          product_slug VARCHAR(100) NOT NULL UNIQUE,
                          title VARCHAR(255) NOT NULL,
                          description TEXT NOT NULL,
                          price DECIMAL(12,2) NOT NULL DEFAULT 0.00 CHECK (price >= 0),
                          condition ProductCondition NOT NULL DEFAULT 'Like New',
                          status ProductStatus NOT NULL DEFAULT 'Available',
                          view_count INTEGER NOT NULL DEFAULT 0,
                          created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create Location table
CREATE TABLE Location (
                          location_id SERIAL PRIMARY KEY,
                          product_id INTEGER REFERENCES Products(product_id) ON DELETE CASCADE,
                          latitude DECIMAL(10,7),
                          longitude DECIMAL(10,7),
                          address TEXT,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- CategoryProductSpecs: now linked to SubCategories (not main Categories)
CREATE TABLE CategoryProductSpecs (
                                      category_product_spec_id SERIAL PRIMARY KEY,
                                      sub_category_id INTEGER NOT NULL REFERENCES SubCategories(sub_category_id) ON DELETE CASCADE,
                                      is_main_spec BOOLEAN DEFAULT FALSE,
                                      spec_type_title VARCHAR(100) NOT NULL,
                                      spec_name VARCHAR(100) NOT NULL,
                                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                      UNIQUE (sub_category_id, spec_name)
);

-- Create ProductSpecs table
CREATE TABLE ProductSpecs (
                              product_spec_id SERIAL PRIMARY KEY,
                              category_product_spec_id INTEGER REFERENCES CategoryProductSpecs(category_product_spec_id) ON DELETE CASCADE,
                              product_id INTEGER REFERENCES Products(product_id) ON DELETE CASCADE,
                              spec_value VARCHAR(100) NOT NULL,
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE ProductSpecs
    ADD CONSTRAINT uk_product_specs_product_id_spec_id
        UNIQUE (product_id, category_product_spec_id);

-- Create ProductImages table
CREATE TABLE ProductImages (
                               image_id SERIAL PRIMARY KEY,
                               product_id INTEGER NOT NULL REFERENCES Products(product_id) ON DELETE CASCADE,
                               image_url VARCHAR(255) NOT NULL,
                               is_primary BOOLEAN DEFAULT FALSE,
                               order_sequence INTEGER DEFAULT 0,
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create Wishlists table
CREATE TABLE Wishlists (
                           wishlist_id SERIAL PRIMARY KEY,
                           user_id INTEGER REFERENCES Users(user_id) ON DELETE CASCADE,
                           product_id INTEGER REFERENCES Products(product_id) ON DELETE CASCADE,
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           UNIQUE (user_id, product_id)
);

-- Create Bookings table
CREATE TABLE Bookings (
                          booking_id SERIAL PRIMARY KEY,
                          user_id INTEGER REFERENCES Users(user_id) ON DELETE CASCADE,
                          product_id INTEGER REFERENCES Products(product_id) ON DELETE SET NULL,
                          order_number VARCHAR(50) NOT NULL UNIQUE,
                          total_amount DECIMAL(12,2) NOT NULL,
                          book_status BookingStatus DEFAULT 'Pending',
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create Payments table
CREATE TABLE Payments (
                          payment_id SERIAL PRIMARY KEY,
                          booking_id INTEGER REFERENCES Bookings(booking_id) ON DELETE CASCADE,
                          transaction_id VARCHAR(100) UNIQUE,
                          payment_method VARCHAR(50) NOT NULL,
                          amount DECIMAL(12,2) NOT NULL,
                          payment_status PaymentStatus NOT NULL DEFAULT 'Pending',
                          payment_date TIMESTAMP,
                          payment_proof TEXT,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create UserBalances table
CREATE TABLE UserBalances (
                              balance_id SERIAL PRIMARY KEY,
                              user_id INTEGER REFERENCES Users(user_id) ON DELETE CASCADE,
                              amount DECIMAL(12,2) NOT NULL,
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create UserBalanceHistory table
CREATE TABLE UserBalanceHistory (
                                    history_id SERIAL PRIMARY KEY,
                                    balance_id INTEGER REFERENCES UserBalances(balance_id) ON DELETE CASCADE,
                                    amount DECIMAL(12,2) NOT NULL,
                                    balance_type BalanceType DEFAULT 'Top Up',
                                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create Messages table
CREATE TABLE Messages (
                          message_id SERIAL PRIMARY KEY,
                          sender_id INTEGER REFERENCES Users(user_id) ON DELETE CASCADE,
                          receiver_id INTEGER REFERENCES Users(user_id) ON DELETE CASCADE,
                          product_id INTEGER REFERENCES Products(product_id) ON DELETE SET NULL,
                          message_text TEXT NOT NULL,
                          is_read BOOLEAN DEFAULT FALSE,
                          sent_at TIMESTAMP DEFAULT NOW()
);

-- Create Banners table
CREATE TABLE Banners (
                         banner_id SERIAL PRIMARY KEY,
                         image_url VARCHAR(255) NOT NULL,
                         sequence INTEGER DEFAULT 0,
                         link VARCHAR(255),
                         title VARCHAR(100),
                         description VARCHAR(255),
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX idx_users_email ON Users(email);
CREATE INDEX idx_products_listing_user ON Products(listing_user_id);
CREATE INDEX idx_products_subcategory ON Products(sub_category_id);  -- updated
CREATE INDEX idx_products_status ON Products(status);
CREATE INDEX idx_bookings_user ON Bookings(user_id);
CREATE INDEX idx_bookings_product ON Bookings(product_id);
CREATE INDEX idx_messages_sender ON Messages(sender_id);
CREATE INDEX idx_messages_receiver ON Messages(receiver_id);
CREATE INDEX idx_wishlists_user ON Wishlists(user_id);
CREATE INDEX idx_wishlists_product ON Wishlists(product_id);
CREATE INDEX idx_subcategories_parent ON SubCategories(parent_category_id);
CREATE INDEX idx_subcategories_slug ON SubCategories(sub_category_slug);