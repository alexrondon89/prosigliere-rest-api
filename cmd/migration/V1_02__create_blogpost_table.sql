CREATE TABLE IF NOT EXISTS blog_post (
    id  UUID    PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT  NOT NULL,
    content TEXT    NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT  NOW()
);
