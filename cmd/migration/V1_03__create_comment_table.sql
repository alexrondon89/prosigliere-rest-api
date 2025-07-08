CREATE TABLE IF NOT EXISTS comment (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    blog_post_id  UUID NOT NULL REFERENCES blog_post(id) ON DELETE CASCADE,
    username      VARCHAR(30) NOT NULL,
    content       VARCHAR(4000) NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_comments_post ON comment (blog_post_id);