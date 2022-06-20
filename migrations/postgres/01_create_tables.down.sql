ALTER TABLE IF EXISTS "article" DROP CONSTRAINT IF EXISTS "article_author_id_fkey";

DROP INDEX IF EXISTS "unique_firstname_lastname_on_author";

DROP TABLE IF EXISTS "author";

DROP INDEX IF EXISTS "unique_title_on_article";

DROP TABLE IF EXISTS "article";

DROP EXTENSION IF EXISTS "uuid-ossp";