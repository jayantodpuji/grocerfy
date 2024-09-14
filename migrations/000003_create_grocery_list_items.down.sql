DROP INDEX idx_grocery_list_items_grocery_list_id if exists;
DROP INDEX idx_grocery_list_items_category if exists;
DROP CONSTRAINT fk_grocery_list_items_grocery_list if exists;

DROP TABLE public.grocery_list_items if exists;