UPDATE posts
SET content = 'The post content was updated!',
    updated_at = NOW(),
WHERE id = 2;