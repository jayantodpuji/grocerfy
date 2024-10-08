CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.grocery_list_items (
  id uuid primary key default uuid_generate_v4(),
  grocery_list_id uuid not null,
  name text not null,
  unit text not null,
  quantity int not null,
  is_purchased boolean not null default false,
  created_at timestamptz not null default current_timestamp,
  updated_at timestamptz not null default current_timestamp,

  CONSTRAINT fk_grocery_list_items_grocery_list FOREIGN KEY (grocery_list_id) REFERENCES public.grocery_lists(id)
);

CREATE INDEX idx_grocery_list_items_grocery_list_id ON public.grocery_list_items (grocery_list_id);
