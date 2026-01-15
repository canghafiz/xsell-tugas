-- ==============================
-- INSERT: MAIN CATEGORIES
-- ==============================

INSERT INTO Categories (category_name, category_slug, description, icon) VALUES
                                                                             ('Electronics', 'electronics', 'Phones, laptops, cameras, and accessories', 'smartphone'),
                                                                             ('Fashion', 'fashion', 'Clothing, shoes, bags, and fashion accessories', 'shirt'),
                                                                             ('Automotive', 'automotive', 'Cars, motorcycles, parts, and vehicle accessories', 'car'),
                                                                             ('Property', 'property', 'Houses, apartments, rentals, and real estate', 'home'),
                                                                             ('Furniture', 'furniture', 'Tables, chairs, beds, and home furnishings', 'bed'),
                                                                             ('Hobbies & Collectibles', 'hobbies_collectibles', 'Toys, vinyl records, collectible cards, and memorabilia', 'gamepad'),
                                                                             ('Books & Stationery', 'books_stationery', 'Textbooks, novels, notebooks, and office supplies', 'book'),
                                                                             ('Sports', 'sports', 'Bikes, fitness equipment, and outdoor gear', 'dumbbell'),
                                                                             ('Kids & Baby', 'kids_baby', 'Baby clothes, strollers, educational toys', 'baby'),
                                                                             ('Services', 'services', 'Freelance work, tutoring, repairs, and local gigs', 'briefcase'),
                                                                             ('Other', 'other', 'Items that don''t fit in other categories', 'package');

-- ==============================
-- INSERT: SUBCATEGORIES
-- ==============================

-- Electronics
INSERT INTO SubCategories (parent_category_id, sub_category_name, sub_category_slug, icon)
SELECT category_id, 'Smartphone', 'smartphone', 'smartphone' FROM Categories WHERE category_slug = 'electronics'
UNION ALL SELECT category_id, 'Laptop', 'laptop', 'laptop' FROM Categories WHERE category_slug = 'electronics'
UNION ALL SELECT category_id, 'Tablet', 'tablet', 'tablet' FROM Categories WHERE category_slug = 'electronics'
UNION ALL SELECT category_id, 'Camera', 'camera', 'camera' FROM Categories WHERE category_slug = 'electronics'
UNION ALL SELECT category_id, 'Headphones & Earphones', 'headphones', 'headphones' FROM Categories WHERE category_slug = 'electronics';

-- Fashion
INSERT INTO SubCategories (parent_category_id, sub_category_name, sub_category_slug, icon)
SELECT category_id, 'Men''s Clothing', 'mens-clothing', 'shirt' FROM Categories WHERE category_slug = 'fashion'
UNION ALL SELECT category_id, 'Women''s Clothing', 'womens-clothing', 'dress' FROM Categories WHERE category_slug = 'fashion'
UNION ALL SELECT category_id, 'Shoes', 'shoes', 'shoe' FROM Categories WHERE category_slug = 'fashion'
UNION ALL SELECT category_id, 'Bags & Wallets', 'bags-wallets', 'bag' FROM Categories WHERE category_slug = 'fashion'
UNION ALL SELECT category_id, 'Watches & Accessories', 'watches-accessories', 'watch' FROM Categories WHERE category_slug = 'fashion';

-- Automotive
INSERT INTO SubCategories (parent_category_id, sub_category_name, sub_category_slug, icon)
SELECT category_id, 'Cars', 'cars', 'car' FROM Categories WHERE category_slug = 'automotive'
UNION ALL SELECT category_id, 'Motorcycles', 'motorcycles', 'motorcycle' FROM Categories WHERE category_slug = 'automotive'
UNION ALL SELECT category_id, 'Car Parts & Accessories', 'car-parts', 'tools' FROM Categories WHERE category_slug = 'automotive'
UNION ALL SELECT category_id, 'Motorcycle Parts', 'motorcycle-parts', 'tools' FROM Categories WHERE category_slug = 'automotive';

-- Furniture
INSERT INTO SubCategories (parent_category_id, sub_category_name, sub_category_slug, icon)
SELECT category_id, 'Living Room', 'living-room', 'sofa' FROM Categories WHERE category_slug = 'furniture'
UNION ALL SELECT category_id, 'Bedroom', 'bedroom', 'bed' FROM Categories WHERE category_slug = 'furniture'
UNION ALL SELECT category_id, 'Dining Room', 'dining-room', 'table' FROM Categories WHERE category_slug = 'furniture'
UNION ALL SELECT category_id, 'Office Furniture', 'office-furniture', 'desk' FROM Categories WHERE category_slug = 'furniture';

-- Sports
INSERT INTO SubCategories (parent_category_id, sub_category_name, sub_category_slug, icon)
SELECT category_id, 'Fitness Equipment', 'fitness-equipment', 'dumbbell' FROM Categories WHERE category_slug = 'sports'
UNION ALL SELECT category_id, 'Bicycles', 'bicycles', 'bicycle' FROM Categories WHERE category_slug = 'sports'
UNION ALL SELECT category_id, 'Outdoor Gear', 'outdoor-gear', 'tent' FROM Categories WHERE category_slug = 'sports';

-- Kids & Baby
INSERT INTO SubCategories (parent_category_id, sub_category_name, sub_category_slug, icon)
SELECT category_id, 'Toys', 'toys', 'toy' FROM Categories WHERE category_slug = 'kids_baby'
UNION ALL SELECT category_id, 'Baby Gear', 'baby-gear', 'stroller' FROM Categories WHERE category_slug = 'kids_baby'
UNION ALL SELECT category_id, 'Kids Clothing', 'kids-clothing', 'tshirt' FROM Categories WHERE category_slug = 'kids_baby';

-- Books & Stationery
INSERT INTO SubCategories (parent_category_id, sub_category_name, sub_category_slug, icon)
SELECT category_id, 'Books', 'books', 'book' FROM Categories WHERE category_slug = 'books_stationery'
UNION ALL SELECT category_id, 'Stationery', 'stationery', 'pen' FROM Categories WHERE category_slug = 'books_stationery'
UNION ALL SELECT category_id, 'Art Supplies', 'art-supplies', 'palette' FROM Categories WHERE category_slug = 'books_stationery';

-- Hobbies & Collectibles
INSERT INTO SubCategories (parent_category_id, sub_category_name, sub_category_slug, icon)
SELECT category_id, 'Collectible Cards', 'collectible-cards', 'cards' FROM Categories WHERE category_slug = 'hobbies_collectibles'
UNION ALL SELECT category_id, 'Vinyl Records', 'vinyl-records', 'vinyl' FROM Categories WHERE category_slug = 'hobbies_collectibles'
UNION ALL SELECT category_id, 'Action Figures', 'action-figures', 'figure' FROM Categories WHERE category_slug = 'hobbies_collectibles';

-- Property
INSERT INTO SubCategories (parent_category_id, sub_category_name, sub_category_slug, icon)
SELECT category_id, 'Houses', 'houses', 'home' FROM Categories WHERE category_slug = 'property'
UNION ALL SELECT category_id, 'Apartments', 'apartments', 'building' FROM Categories WHERE category_slug = 'property'
UNION ALL SELECT category_id, 'Land', 'land', 'map' FROM Categories WHERE category_slug = 'property'
UNION ALL SELECT category_id, 'Commercial Property', 'commercial-property', 'store' FROM Categories WHERE category_slug = 'property';

-- Services
INSERT INTO SubCategories (parent_category_id, sub_category_name, sub_category_slug, icon)
SELECT category_id, 'Tutoring', 'tutoring', 'book-open' FROM Categories WHERE category_slug = 'services'
UNION ALL SELECT category_id, 'Repair Services', 'repair-services', 'wrench' FROM Categories WHERE category_slug = 'services'
UNION ALL SELECT category_id, 'Freelance', 'freelance', 'briefcase' FROM Categories WHERE category_slug = 'services'
UNION ALL SELECT category_id, 'Beauty Services', 'beauty-services', 'sparkles' FROM Categories WHERE category_slug = 'services';

-- Other
INSERT INTO SubCategories (parent_category_id, sub_category_name, sub_category_slug, icon)
SELECT category_id, 'Miscellaneous', 'miscellaneous', 'package' FROM Categories WHERE category_slug = 'other';

-- ==============================
-- INSERT: PRODUCT SPECS
-- ==============================

-- 📱 Smartphone
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'smartphone'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Model' FROM SubCategories WHERE sub_category_slug = 'smartphone'
UNION ALL SELECT sub_category_id, TRUE, 'Display', 'Screen Size (inches)' FROM SubCategories WHERE sub_category_slug = 'smartphone'
UNION ALL SELECT sub_category_id, TRUE, 'Camera', 'Rear Camera (MP)' FROM SubCategories WHERE sub_category_slug = 'smartphone'
UNION ALL SELECT sub_category_id, TRUE, 'Camera', 'Front Camera (MP)' FROM SubCategories WHERE sub_category_slug = 'smartphone'
UNION ALL SELECT sub_category_id, TRUE, 'Memory', 'RAM (GB)' FROM SubCategories WHERE sub_category_slug = 'smartphone'
UNION ALL SELECT sub_category_id, TRUE, 'Storage', 'Internal Storage (GB)' FROM SubCategories WHERE sub_category_slug = 'smartphone'
UNION ALL SELECT sub_category_id, FALSE, 'Battery', 'Battery Capacity (mAh)' FROM SubCategories WHERE sub_category_slug = 'smartphone'
UNION ALL SELECT sub_category_id, FALSE, 'Connectivity', '5G Support' FROM SubCategories WHERE sub_category_slug = 'smartphone'
UNION ALL SELECT sub_category_id, FALSE, 'Physical', 'Color' FROM SubCategories WHERE sub_category_slug = 'smartphone';

-- 💻 Laptop
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'laptop'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Model' FROM SubCategories WHERE sub_category_slug = 'laptop'
UNION ALL SELECT sub_category_id, TRUE, 'Processor', 'Processor Brand' FROM SubCategories WHERE sub_category_slug = 'laptop'
UNION ALL SELECT sub_category_id, TRUE, 'Processor', 'Processor Model' FROM SubCategories WHERE sub_category_slug = 'laptop'
UNION ALL SELECT sub_category_id, TRUE, 'Memory', 'RAM (GB)' FROM SubCategories WHERE sub_category_slug = 'laptop'
UNION ALL SELECT sub_category_id, TRUE, 'Storage', 'SSD Capacity (GB)' FROM SubCategories WHERE sub_category_slug = 'laptop'
UNION ALL SELECT sub_category_id, TRUE, 'Display', 'Screen Size (inches)' FROM SubCategories WHERE sub_category_slug = 'laptop'
UNION ALL SELECT sub_category_id, FALSE, 'Graphics', 'Dedicated GPU' FROM SubCategories WHERE sub_category_slug = 'laptop'
UNION ALL SELECT sub_category_id, FALSE, 'Ports', 'USB-C Ports' FROM SubCategories WHERE sub_category_slug = 'laptop'
UNION ALL SELECT sub_category_id, FALSE, 'Physical', 'Weight (kg)' FROM SubCategories WHERE sub_category_slug = 'laptop';

-- 📱 Tablet
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'tablet'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Model' FROM SubCategories WHERE sub_category_slug = 'tablet'
UNION ALL SELECT sub_category_id, TRUE, 'Display', 'Screen Size (inches)' FROM SubCategories WHERE sub_category_slug = 'tablet'
UNION ALL SELECT sub_category_id, TRUE, 'Processor', 'Processor Model' FROM SubCategories WHERE sub_category_slug = 'tablet'
UNION ALL SELECT sub_category_id, TRUE, 'Memory', 'RAM (GB)' FROM SubCategories WHERE sub_category_slug = 'tablet'
UNION ALL SELECT sub_category_id, TRUE, 'Storage', 'Internal Storage (GB)' FROM SubCategories WHERE sub_category_slug = 'tablet'
UNION ALL SELECT sub_category_id, FALSE, 'Connectivity', 'Wi-Fi Only' FROM SubCategories WHERE sub_category_slug = 'tablet'
UNION ALL SELECT sub_category_id, FALSE, 'Physical', 'Color' FROM SubCategories WHERE sub_category_slug = 'tablet';

-- 📷 Camera
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'camera'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Model' FROM SubCategories WHERE sub_category_slug = 'camera'
UNION ALL SELECT sub_category_id, TRUE, 'Camera', 'Rear Camera (MP)' FROM SubCategories WHERE sub_category_slug = 'camera'
UNION ALL SELECT sub_category_id, FALSE, 'General', 'Release Year' FROM SubCategories WHERE sub_category_slug = 'camera'
UNION ALL SELECT sub_category_id, FALSE, 'Type', 'Camera Type' FROM SubCategories WHERE sub_category_slug = 'camera';

-- 🎧 Headphones
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'headphones'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Model' FROM SubCategories WHERE sub_category_slug = 'headphones'
UNION ALL SELECT sub_category_id, TRUE, 'Connectivity', 'Bluetooth Version' FROM SubCategories WHERE sub_category_slug = 'headphones'
UNION ALL SELECT sub_category_id, FALSE, 'Features', 'Noise Cancellation' FROM SubCategories WHERE sub_category_slug = 'headphones'
UNION ALL SELECT sub_category_id, FALSE, 'Physical', 'Color' FROM SubCategories WHERE sub_category_slug = 'headphones';

-- 👕 Men's Clothing
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'mens-clothing'
UNION ALL SELECT sub_category_id, TRUE, 'Size & Fit', 'Size (S/M/L/XL)' FROM SubCategories WHERE sub_category_slug = 'mens-clothing'
UNION ALL SELECT sub_category_id, TRUE, 'Material', 'Fabric Material' FROM SubCategories WHERE sub_category_slug = 'mens-clothing'
UNION ALL SELECT sub_category_id, FALSE, 'Details', 'Color' FROM SubCategories WHERE sub_category_slug = 'mens-clothing'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Stains or Damage' FROM SubCategories WHERE sub_category_slug = 'mens-clothing';

-- 👗 Women's Clothing
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'womens-clothing'
UNION ALL SELECT sub_category_id, TRUE, 'Size & Fit', 'Size (S/M/L/XL)' FROM SubCategories WHERE sub_category_slug = 'womens-clothing'
UNION ALL SELECT sub_category_id, TRUE, 'Material', 'Fabric Material' FROM SubCategories WHERE sub_category_slug = 'womens-clothing'
UNION ALL SELECT sub_category_id, FALSE, 'Details', 'Color' FROM SubCategories WHERE sub_category_slug = 'womens-clothing'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Stains or Damage' FROM SubCategories WHERE sub_category_slug = 'womens-clothing';

-- 👠 Shoes
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Product Type' FROM SubCategories WHERE sub_category_slug = 'shoes'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'shoes'
UNION ALL SELECT sub_category_id, TRUE, 'Size', 'Size (S/M/L/XL)' FROM SubCategories WHERE sub_category_slug = 'shoes'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Gender' FROM SubCategories WHERE sub_category_slug = 'shoes'
UNION ALL SELECT sub_category_id, FALSE, 'Material', 'Material' FROM SubCategories WHERE sub_category_slug = 'shoes'
UNION ALL SELECT sub_category_id, FALSE, 'Physical', 'Color' FROM SubCategories WHERE sub_category_slug = 'shoes';

-- 👜 Bags & Wallets
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Product Type' FROM SubCategories WHERE sub_category_slug = 'bags-wallets'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'bags-wallets'
UNION ALL SELECT sub_category_id, TRUE, 'Material', 'Material' FROM SubCategories WHERE sub_category_slug = 'bags-wallets'
UNION ALL SELECT sub_category_id, TRUE, 'Physical', 'Color' FROM SubCategories WHERE sub_category_slug = 'bags-wallets'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Gender' FROM SubCategories WHERE sub_category_slug = 'bags-wallets'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Authenticity' FROM SubCategories WHERE sub_category_slug = 'bags-wallets';

-- ⌚ Watches & Accessories
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Product Type' FROM SubCategories WHERE sub_category_slug = 'watches-accessories'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'watches-accessories'
UNION ALL SELECT sub_category_id, TRUE, 'Material', 'Material' FROM SubCategories WHERE sub_category_slug = 'watches-accessories'
UNION ALL SELECT sub_category_id, TRUE, 'Physical', 'Color' FROM SubCategories WHERE sub_category_slug = 'watches-accessories'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Gender' FROM SubCategories WHERE sub_category_slug = 'watches-accessories'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Original Box' FROM SubCategories WHERE sub_category_slug = 'watches-accessories';

-- 🚗 Cars
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'Vehicle', 'Brand' FROM SubCategories WHERE sub_category_slug = 'cars'
UNION ALL SELECT sub_category_id, TRUE, 'Vehicle', 'Model' FROM SubCategories WHERE sub_category_slug = 'cars'
UNION ALL SELECT sub_category_id, TRUE, 'Vehicle', 'Year' FROM SubCategories WHERE sub_category_slug = 'cars'
UNION ALL SELECT sub_category_id, TRUE, 'Vehicle', 'Mileage (km)' FROM SubCategories WHERE sub_category_slug = 'cars'
UNION ALL SELECT sub_category_id, TRUE, 'Engine', 'Fuel Type' FROM SubCategories WHERE sub_category_slug = 'cars'
UNION ALL SELECT sub_category_id, TRUE, 'Transmission', 'Transmission Type' FROM SubCategories WHERE sub_category_slug = 'cars'
UNION ALL SELECT sub_category_id, FALSE, 'Legal', 'Registration Status' FROM SubCategories WHERE sub_category_slug = 'cars'
UNION ALL SELECT sub_category_id, FALSE, 'Features', 'Sunroof' FROM SubCategories WHERE sub_category_slug = 'cars'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Accident History' FROM SubCategories WHERE sub_category_slug = 'cars';

-- 🏍️ Motorcycles
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'Vehicle', 'Brand' FROM SubCategories WHERE sub_category_slug = 'motorcycles'
UNION ALL SELECT sub_category_id, TRUE, 'Vehicle', 'Model' FROM SubCategories WHERE sub_category_slug = 'motorcycles'
UNION ALL SELECT sub_category_id, TRUE, 'Vehicle', 'Year' FROM SubCategories WHERE sub_category_slug = 'motorcycles'
UNION ALL SELECT sub_category_id, TRUE, 'Vehicle', 'Mileage (km)' FROM SubCategories WHERE sub_category_slug = 'motorcycles'
UNION ALL SELECT sub_category_id, TRUE, 'Engine', 'Engine (cc)' FROM SubCategories WHERE sub_category_slug = 'motorcycles'
UNION ALL SELECT sub_category_id, TRUE, 'Engine', 'Fuel Type' FROM SubCategories WHERE sub_category_slug = 'motorcycles'
UNION ALL SELECT sub_category_id, TRUE, 'Transmission', 'Transmission Type' FROM SubCategories WHERE sub_category_slug = 'motorcycles'
UNION ALL SELECT sub_category_id, FALSE, 'Legal', 'Registration Status' FROM SubCategories WHERE sub_category_slug = 'motorcycles'
UNION ALL SELECT sub_category_id, FALSE, 'Physical', 'Color' FROM SubCategories WHERE sub_category_slug = 'motorcycles';

-- 🔧 Car Parts
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'car-parts'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Model' FROM SubCategories WHERE sub_category_slug = 'car-parts'
UNION ALL SELECT sub_category_id, TRUE, 'Specifications', 'Variant' FROM SubCategories WHERE sub_category_slug = 'car-parts'
UNION ALL SELECT sub_category_id, FALSE, 'Compatibility', 'Compatible Brands' FROM SubCategories WHERE sub_category_slug = 'car-parts'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Warranty' FROM SubCategories WHERE sub_category_slug = 'car-parts';

-- 🏍️ Motorcycle Parts
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'motorcycle-parts'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Model' FROM SubCategories WHERE sub_category_slug = 'motorcycle-parts'
UNION ALL SELECT sub_category_id, TRUE, 'Specifications', 'Variant' FROM SubCategories WHERE sub_category_slug = 'motorcycle-parts'
UNION ALL SELECT sub_category_id, FALSE, 'General', 'Year' FROM SubCategories WHERE sub_category_slug = 'motorcycle-parts'
UNION ALL SELECT sub_category_id, FALSE, 'Physical', 'Color' FROM SubCategories WHERE sub_category_slug = 'motorcycle-parts'
UNION ALL SELECT sub_category_id, FALSE, 'Safety', 'Certification' FROM SubCategories WHERE sub_category_slug = 'motorcycle-parts';

-- 🛋️ Living Room
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Furniture Type' FROM SubCategories WHERE sub_category_slug = 'living-room'
UNION ALL SELECT sub_category_id, TRUE, 'Material', 'Frame Material' FROM SubCategories WHERE sub_category_slug = 'living-room'
UNION ALL SELECT sub_category_id, TRUE, 'Material', 'Primary Material' FROM SubCategories WHERE sub_category_slug = 'living-room'
UNION ALL SELECT sub_category_id, TRUE, 'Dimensions', 'Width (cm)' FROM SubCategories WHERE sub_category_slug = 'living-room'
UNION ALL SELECT sub_category_id, TRUE, 'Dimensions', 'Height (cm)' FROM SubCategories WHERE sub_category_slug = 'living-room'
UNION ALL SELECT sub_category_id, TRUE, 'Dimensions', 'Depth (cm)' FROM SubCategories WHERE sub_category_slug = 'living-room'
UNION ALL SELECT sub_category_id, FALSE, 'Style', 'Design Style' FROM SubCategories WHERE sub_category_slug = 'living-room'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Scratches or Wear' FROM SubCategories WHERE sub_category_slug = 'living-room';

-- 🛏️ Bedroom
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Furniture Type' FROM SubCategories WHERE sub_category_slug = 'bedroom'
UNION ALL SELECT sub_category_id, TRUE, 'Material', 'Primary Material' FROM SubCategories WHERE sub_category_slug = 'bedroom'
UNION ALL SELECT sub_category_id, TRUE, 'Dimensions', 'Height (cm)' FROM SubCategories WHERE sub_category_slug = 'bedroom'
UNION ALL SELECT sub_category_id, TRUE, 'Dimensions', 'Width (cm)' FROM SubCategories WHERE sub_category_slug = 'bedroom'
UNION ALL SELECT sub_category_id, FALSE, 'Dimensions', 'Depth (cm)' FROM SubCategories WHERE sub_category_slug = 'bedroom'
UNION ALL SELECT sub_category_id, FALSE, 'Features', 'Has Mirror' FROM SubCategories WHERE sub_category_slug = 'bedroom'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Scratches or Wear' FROM SubCategories WHERE sub_category_slug = 'bedroom';

-- 🍽️ Dining Room
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Furniture Type' FROM SubCategories WHERE sub_category_slug = 'dining-room'
UNION ALL SELECT sub_category_id, TRUE, 'Material', 'Primary Material' FROM SubCategories WHERE sub_category_slug = 'dining-room'
UNION ALL SELECT sub_category_id, TRUE, 'Dimensions', 'Width (cm)' FROM SubCategories WHERE sub_category_slug = 'dining-room'
UNION ALL SELECT sub_category_id, TRUE, 'Dimensions', 'Depth (cm)' FROM SubCategories WHERE sub_category_slug = 'dining-room'
UNION ALL SELECT sub_category_id, FALSE, 'Dimensions', 'Height (cm)' FROM SubCategories WHERE sub_category_slug = 'dining-room'
UNION ALL SELECT sub_category_id, FALSE, 'Features', 'Seating Capacity' FROM SubCategories WHERE sub_category_slug = 'dining-room'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Includes Chairs' FROM SubCategories WHERE sub_category_slug = 'dining-room';

-- 🪑 Office Furniture
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Furniture Type' FROM SubCategories WHERE sub_category_slug = 'office-furniture'
UNION ALL SELECT sub_category_id, TRUE, 'Material', 'Primary Material' FROM SubCategories WHERE sub_category_slug = 'office-furniture'
UNION ALL SELECT sub_category_id, TRUE, 'Dimensions', 'Width (cm)' FROM SubCategories WHERE sub_category_slug = 'office-furniture'
UNION ALL SELECT sub_category_id, TRUE, 'Dimensions', 'Depth (cm)' FROM SubCategories WHERE sub_category_slug = 'office-furniture'
UNION ALL SELECT sub_category_id, FALSE, 'Dimensions', 'Height (cm)' FROM SubCategories WHERE sub_category_slug = 'office-furniture'
UNION ALL SELECT sub_category_id, FALSE, 'Material', 'Upholstery Material' FROM SubCategories WHERE sub_category_slug = 'office-furniture'
UNION ALL SELECT sub_category_id, FALSE, 'Physical', 'Primary Color' FROM SubCategories WHERE sub_category_slug = 'office-furniture'
UNION ALL SELECT sub_category_id, FALSE, 'Features', 'Adjustable Height' FROM SubCategories WHERE sub_category_slug = 'office-furniture'
UNION ALL SELECT sub_category_id, FALSE, 'Features', 'Wheels/Casters' FROM SubCategories WHERE sub_category_slug = 'office-furniture';

-- 🏋️ Fitness Equipment
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Equipment Type' FROM SubCategories WHERE sub_category_slug = 'fitness-equipment'
UNION ALL SELECT sub_category_id, TRUE, 'Brand', 'Brand' FROM SubCategories WHERE sub_category_slug = 'fitness-equipment'
UNION ALL SELECT sub_category_id, TRUE, 'Specs', 'Max Weight Capacity (kg)' FROM SubCategories WHERE sub_category_slug = 'fitness-equipment'
UNION ALL SELECT sub_category_id, TRUE, 'Specs', 'Dimensions (L x W x H cm)' FROM SubCategories WHERE sub_category_slug = 'fitness-equipment'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Used Duration (months)' FROM SubCategories WHERE sub_category_slug = 'fitness-equipment';

-- 🚴 Bicycles
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'bicycles'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Model' FROM SubCategories WHERE sub_category_slug = 'bicycles'
UNION ALL SELECT sub_category_id, TRUE, 'Type', 'Bicycle Type' FROM SubCategories WHERE sub_category_slug = 'bicycles'
UNION ALL SELECT sub_category_id, TRUE, 'Specs', 'Frame Size' FROM SubCategories WHERE sub_category_slug = 'bicycles'
UNION ALL SELECT sub_category_id, FALSE, 'Material', 'Frame Material' FROM SubCategories WHERE sub_category_slug = 'bicycles'
UNION ALL SELECT sub_category_id, FALSE, 'Specs', 'Number of Gears' FROM SubCategories WHERE sub_category_slug = 'bicycles';

-- ⛺ Outdoor Gear
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Product Type' FROM SubCategories WHERE sub_category_slug = 'outdoor-gear'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'outdoor-gear'
UNION ALL SELECT sub_category_id, TRUE, 'Material', 'Material' FROM SubCategories WHERE sub_category_slug = 'outdoor-gear'
UNION ALL SELECT sub_category_id, FALSE, 'Specs', 'Capacity' FROM SubCategories WHERE sub_category_slug = 'outdoor-gear'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Usage Count' FROM SubCategories WHERE sub_category_slug = 'outdoor-gear';

-- 🧸 Toys
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Toy Type' FROM SubCategories WHERE sub_category_slug = 'toys'
UNION ALL SELECT sub_category_id, TRUE, 'Safety', 'Recommended Age (years)' FROM SubCategories WHERE sub_category_slug = 'toys'
UNION ALL SELECT sub_category_id, TRUE, 'Safety', 'Certified Non-Toxic' FROM SubCategories WHERE sub_category_slug = 'toys'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Missing Parts' FROM SubCategories WHERE sub_category_slug = 'toys';

-- 👶 Baby Gear
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Product Type' FROM SubCategories WHERE sub_category_slug = 'baby-gear'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'baby-gear'
UNION ALL SELECT sub_category_id, TRUE, 'Safety', 'Age Range' FROM SubCategories WHERE sub_category_slug = 'baby-gear'
UNION ALL SELECT sub_category_id, FALSE, 'Safety', 'Safety Certification' FROM SubCategories WHERE sub_category_slug = 'baby-gear'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Original Packaging' FROM SubCategories WHERE sub_category_slug = 'baby-gear';

-- 👕 Kids Clothing
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'kids-clothing'
UNION ALL SELECT sub_category_id, TRUE, 'Size & Fit', 'Size/Age' FROM SubCategories WHERE sub_category_slug = 'kids-clothing'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Gender' FROM SubCategories WHERE sub_category_slug = 'kids-clothing'
UNION ALL SELECT sub_category_id, FALSE, 'Material', 'Fabric Material' FROM SubCategories WHERE sub_category_slug = 'kids-clothing'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Stains or Damage' FROM SubCategories WHERE sub_category_slug = 'kids-clothing';

-- 📚 Books
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'Book Info', 'Title' FROM SubCategories WHERE sub_category_slug = 'books'
UNION ALL SELECT sub_category_id, TRUE, 'Book Info', 'Author' FROM SubCategories WHERE sub_category_slug = 'books'
UNION ALL SELECT sub_category_id, TRUE, 'Book Info', 'ISBN' FROM SubCategories WHERE sub_category_slug = 'books'
UNION ALL SELECT sub_category_id, TRUE, 'Book Info', 'Language' FROM SubCategories WHERE sub_category_slug = 'books'
UNION ALL SELECT sub_category_id, TRUE, 'Book Info', 'Pages' FROM SubCategories WHERE sub_category_slug = 'books'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Cover Type (Hard/Paper)' FROM SubCategories WHERE sub_category_slug = 'books'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Notes or Highlights' FROM SubCategories WHERE sub_category_slug = 'books';

-- ✏️ Stationery
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Product Type' FROM SubCategories WHERE sub_category_slug = 'stationery'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'stationery'
UNION ALL SELECT sub_category_id, TRUE, 'Specs', 'Quantity' FROM SubCategories WHERE sub_category_slug = 'stationery'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Sealed/Opened' FROM SubCategories WHERE sub_category_slug = 'stationery';

-- 🎨 Art Supplies
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Product Type' FROM SubCategories WHERE sub_category_slug = 'art-supplies'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'art-supplies'
UNION ALL SELECT sub_category_id, TRUE, 'Specs', 'Set Size' FROM SubCategories WHERE sub_category_slug = 'art-supplies'
UNION ALL SELECT sub_category_id, FALSE, 'Quality', 'Professional/Student Grade' FROM SubCategories WHERE sub_category_slug = 'art-supplies'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Sealed/Opened' FROM SubCategories WHERE sub_category_slug = 'art-supplies';

-- 🃏 Collectible Cards
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Card Game' FROM SubCategories WHERE sub_category_slug = 'collectible-cards'
UNION ALL SELECT sub_category_id, TRUE, 'Card Info', 'Card Name' FROM SubCategories WHERE sub_category_slug = 'collectible-cards'
UNION ALL SELECT sub_category_id, TRUE, 'Card Info', 'Rarity' FROM SubCategories WHERE sub_category_slug = 'collectible-cards'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Card Condition' FROM SubCategories WHERE sub_category_slug = 'collectible-cards'
UNION ALL SELECT sub_category_id, FALSE, 'Grading', 'Graded (Yes/No)' FROM SubCategories WHERE sub_category_slug = 'collectible-cards';

-- 💿 Vinyl Records
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'Album Info', 'Artist' FROM SubCategories WHERE sub_category_slug = 'vinyl-records'
UNION ALL SELECT sub_category_id, TRUE, 'Album Info', 'Album Title' FROM SubCategories WHERE sub_category_slug = 'vinyl-records'
UNION ALL SELECT sub_category_id, TRUE, 'Album Info', 'Release Year' FROM SubCategories WHERE sub_category_slug = 'vinyl-records'
UNION ALL SELECT sub_category_id, FALSE, 'Format', 'Size (12"/7")' FROM SubCategories WHERE sub_category_slug = 'vinyl-records'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Vinyl Condition' FROM SubCategories WHERE sub_category_slug = 'vinyl-records'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Sleeve Condition' FROM SubCategories WHERE sub_category_slug = 'vinyl-records';

-- 🦸 Action Figures
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Character/Series' FROM SubCategories WHERE sub_category_slug = 'action-figures'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'action-figures'
UNION ALL SELECT sub_category_id, TRUE, 'Specs', 'Scale/Size' FROM SubCategories WHERE sub_category_slug = 'action-figures'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Box Condition' FROM SubCategories WHERE sub_category_slug = 'action-figures'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Complete Accessories' FROM SubCategories WHERE sub_category_slug = 'action-figures';

-- 🏠 Houses
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'Property', 'Property Type' FROM SubCategories WHERE sub_category_slug = 'houses'
UNION ALL SELECT sub_category_id, TRUE, 'Size', 'Land Area (sqm)' FROM SubCategories WHERE sub_category_slug = 'houses'
UNION ALL SELECT sub_category_id, TRUE, 'Size', 'Building Area (sqm)' FROM SubCategories WHERE sub_category_slug = 'houses'
UNION ALL SELECT sub_category_id, TRUE, 'Layout', 'Bedrooms' FROM SubCategories WHERE sub_category_slug = 'houses'
UNION ALL SELECT sub_category_id, TRUE, 'Layout', 'Bathrooms' FROM SubCategories WHERE sub_category_slug = 'houses'
UNION ALL SELECT sub_category_id, FALSE, 'Features', 'Garage' FROM SubCategories WHERE sub_category_slug = 'houses'
UNION ALL SELECT sub_category_id, FALSE, 'Legal', 'Certificate Type' FROM SubCategories WHERE sub_category_slug = 'houses';

-- 🏢 Apartments
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'Property', 'Apartment Name' FROM SubCategories WHERE sub_category_slug = 'apartments'
UNION ALL SELECT sub_category_id, TRUE, 'Size', 'Area (sqm)' FROM SubCategories WHERE sub_category_slug = 'apartments'
UNION ALL SELECT sub_category_id, TRUE, 'Layout', 'Bedrooms' FROM SubCategories WHERE sub_category_slug = 'apartments'
UNION ALL SELECT sub_category_id, TRUE, 'Layout', 'Bathrooms' FROM SubCategories WHERE sub_category_slug = 'apartments'
UNION ALL SELECT sub_category_id, TRUE, 'Location', 'Floor' FROM SubCategories WHERE sub_category_slug = 'apartments'
UNION ALL SELECT sub_category_id, FALSE, 'Features', 'Furnished' FROM SubCategories WHERE sub_category_slug = 'apartments'
UNION ALL SELECT sub_category_id, FALSE, 'Legal', 'Certificate Type' FROM SubCategories WHERE sub_category_slug = 'apartments';

-- 🗺️ Land
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'Property', 'Land Type' FROM SubCategories WHERE sub_category_slug = 'land'
UNION ALL SELECT sub_category_id, TRUE, 'Size', 'Area (sqm)' FROM SubCategories WHERE sub_category_slug = 'land'
UNION ALL SELECT sub_category_id, TRUE, 'Legal', 'Certificate Type' FROM SubCategories WHERE sub_category_slug = 'land'
UNION ALL SELECT sub_category_id, FALSE, 'Location', 'Zoning' FROM SubCategories WHERE sub_category_slug = 'land'
UNION ALL SELECT sub_category_id, FALSE, 'Access', 'Road Access' FROM SubCategories WHERE sub_category_slug = 'land';

-- 🏪 Commercial Property
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'Property', 'Property Type' FROM SubCategories WHERE sub_category_slug = 'commercial-property'
UNION ALL SELECT sub_category_id, TRUE, 'Size', 'Area (sqm)' FROM SubCategories WHERE sub_category_slug = 'commercial-property'
UNION ALL SELECT sub_category_id, TRUE, 'Location', 'Location Type' FROM SubCategories WHERE sub_category_slug = 'commercial-property'
UNION ALL SELECT sub_category_id, FALSE, 'Features', 'Parking Spaces' FROM SubCategories WHERE sub_category_slug = 'commercial-property'
UNION ALL SELECT sub_category_id, FALSE, 'Legal', 'Certificate Type' FROM SubCategories WHERE sub_category_slug = 'commercial-property';

-- 📖 Tutoring
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'Service', 'Subject' FROM SubCategories WHERE sub_category_slug = 'tutoring'
UNION ALL SELECT sub_category_id, TRUE, 'Service', 'Education Level' FROM SubCategories WHERE sub_category_slug = 'tutoring'
UNION ALL SELECT sub_category_id, TRUE, 'Service', 'Session Type' FROM SubCategories WHERE sub_category_slug = 'tutoring'
UNION ALL SELECT sub_category_id, FALSE, 'Rate', 'Price per Hour' FROM SubCategories WHERE sub_category_slug = 'tutoring';

-- 🔧 Repair Services
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'Service', 'Service Type' FROM SubCategories WHERE sub_category_slug = 'repair-services'
UNION ALL SELECT sub_category_id, TRUE, 'Service', 'Specialization' FROM SubCategories WHERE sub_category_slug = 'repair-services'
UNION ALL SELECT sub_category_id, FALSE, 'Location', 'On-site Available' FROM SubCategories WHERE sub_category_slug = 'repair-services'
UNION ALL SELECT sub_category_id, FALSE, 'Rate', 'Starting Price' FROM SubCategories WHERE sub_category_slug = 'repair-services';

-- 💼 Freelance
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'Service', 'Service Category' FROM SubCategories WHERE sub_category_slug = 'freelance'
UNION ALL SELECT sub_category_id, TRUE, 'Service', 'Expertise Level' FROM SubCategories WHERE sub_category_slug = 'freelance'
UNION ALL SELECT sub_category_id, FALSE, 'Rate', 'Rate Type' FROM SubCategories WHERE sub_category_slug = 'freelance'
UNION ALL SELECT sub_category_id, FALSE, 'Availability', 'Response Time' FROM SubCategories WHERE sub_category_slug = 'freelance';

-- ✨ Beauty Services
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'Service', 'Service Type' FROM SubCategories WHERE sub_category_slug = 'beauty-services'
UNION ALL SELECT sub_category_id, TRUE, 'Service', 'Service Category' FROM SubCategories WHERE sub_category_slug = 'beauty-services'
UNION ALL SELECT sub_category_id, FALSE, 'Location', 'Home Service Available' FROM SubCategories WHERE sub_category_slug = 'beauty-services'
UNION ALL SELECT sub_category_id, FALSE, 'Rate', 'Starting Price' FROM SubCategories WHERE sub_category_slug = 'beauty-services';

-- 📦 Miscellaneous (Other)
INSERT INTO CategoryProductSpecs (sub_category_id, is_main_spec, spec_type_title, spec_name)
SELECT sub_category_id, TRUE, 'General', 'Product Type' FROM SubCategories WHERE sub_category_slug = 'miscellaneous'
UNION ALL SELECT sub_category_id, TRUE, 'General', 'Brand' FROM SubCategories WHERE sub_category_slug = 'miscellaneous'
UNION ALL SELECT sub_category_id, FALSE, 'Condition', 'Condition' FROM SubCategories WHERE sub_category_slug = 'miscellaneous'
UNION ALL SELECT sub_category_id, FALSE, 'Details', 'Description' FROM SubCategories WHERE sub_category_slug = 'miscellaneous';