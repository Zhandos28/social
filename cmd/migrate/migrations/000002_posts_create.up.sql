CREATE TABLE IF NOT EXISTS posts(
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    user_id bigint NOT NULL,
    content text NOT NULL,
    tags VARCHAR(100) [],
    created_at timestamp(0) with time zone NOT NULL DEFAULT now(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT now()
);

ALTER TABLE posts
    ADD CONSTRAINT fk_posts_user_id FOREIGN KEY (user_id) REFERENCES users(id);