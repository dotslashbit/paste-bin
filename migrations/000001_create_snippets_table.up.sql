CREATE TABLE IF NOT EXISTS snippets (
    id bigserial PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    expires_at timestamp(0) with time zone NOT NULL DEFAULT NOW() + INTERVAL '1 day'
);
