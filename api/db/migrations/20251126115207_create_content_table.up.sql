CREATE TABLE content_sections (
                                  section_id    SERIAL PRIMARY KEY,
                                  section_key   VARCHAR(50)  NOT NULL UNIQUE,  -- "best_selling", "featured", dll
                                  title         VARCHAR(100) NOT NULL,
                                  url           VARCHAR(100) NOT NULL,
                                  subtitle      VARCHAR(255),
                                  section_type  VARCHAR(30)  NOT NULL,        -- "dynamic", "fixed", "predefined"
                                  config        JSONB        NOT NULL,         -- structured config (BUKAN raw SQL)
                                  is_active     BOOLEAN      DEFAULT true,
                                  sort_order    INTEGER      DEFAULT 0,
                                  created_at    TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
                                  updated_at    TIMESTAMP    DEFAULT CURRENT_TIMESTAMP
);

-- Index untuk performa
CREATE UNIQUE INDEX idx_content_sections_key ON content_sections(section_key);
CREATE INDEX idx_content_sections_active ON content_sections(is_active);

CREATE TABLE page_layouts (
                              layout_id     SERIAL PRIMARY KEY,
                              page_key      VARCHAR(50) NOT NULL,          -- "home", "category_electronics", "product_detail", dll
                              section_id    INTEGER NOT NULL REFERENCES content_sections(section_id) ON DELETE CASCADE,
                              sort_order    INTEGER NOT NULL DEFAULT 0,    -- urutan tampil di halaman
                              is_active     BOOLEAN DEFAULT true,
                              created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

                              UNIQUE (page_key, section_id)  -- hindari duplikat section di halaman yang sama
);

-- Index untuk performa
CREATE INDEX idx_page_layouts_page_key ON page_layouts(page_key);
CREATE INDEX idx_page_layouts_active ON page_layouts(is_active);
