ALTER TABLE leathergoods ADD CONSTRAINT leathergoods_name_check CHECK (name <> '');
ALTER TABLE leathergoods ADD CONSTRAINT name_length_check CHECK (LENGTH(name) <= 500);
ALTER TABLE leathergoods ADD CONSTRAINT leathergoods_price_check CHECK (price <> 0);
ALTER TABLE leathergoods ADD CONSTRAINT leathergoods_type_check CHECK (type <> '');
ALTER TABLE leathergoods ADD CONSTRAINT leathergoods_leather_type_check CHECK (leather_type <> '');
ALTER TABLE leathergoods ADD CONSTRAINT leathergoods_color_check CHECK (color <> '');

