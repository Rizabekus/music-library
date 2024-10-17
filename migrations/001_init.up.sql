CREATE TABLE IF NOT EXISTS musiclib (
    id SERIAL PRIMARY KEY,
    "group" VARCHAR(255) NOT NULL,
    song VARCHAR(255) NOT NULL,
    releasedate VARCHAR(255) NOT NULL,
    "text" TEXT NOT NULL,
    link VARCHAR(512) NOT NULL    
);
