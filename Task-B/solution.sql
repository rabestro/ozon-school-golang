SELECT goods.id, goods.name
FROM goods
JOIN tags_goods
  ON goods.id = tags_goods.goods_id
GROUP BY goods.id
HAVING COUNT(tag_id) = (SELECT COUNT(*) AS tags_count FROM tags);
