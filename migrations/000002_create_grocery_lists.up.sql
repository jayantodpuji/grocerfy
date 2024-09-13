CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.grocery_lists (
  id uuid primary key default uuid_generate_v4(),
  user_id uuid not null,
  name text not null,
  description text,
  created_at timestamptz not null default current_timestamp,
  updated_at timestamptz not null default current_timestamp,
  deleted_at timestamptz null,
  CONSTRAINT fk_grocery_lists_users FOREIGN KEY (user_id) REFERENCES public.users (id)
);

CREATE INDEX idx_grocery_lists_user_id ON public.grocery_lists (user_id);