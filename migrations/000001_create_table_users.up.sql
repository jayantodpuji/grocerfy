CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.users (
  id uuid primary key default uuid_generate_v4(),
  email text not null unique,
  password_hash text not null,
  name text,
  created_at timestamptz not null default current_timestamp,
  updated_at timestamptz not null default current_timestamp,
  deleted_at timestamptz null
);

CREATE INDEX idx_users_email ON public.users (email);
