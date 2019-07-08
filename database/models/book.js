'use strict';
module.exports = (sequelize, DataTypes) => {
  const book = sequelize.define('book', {
    cover: DataTypes.STRING,
    author: DataTypes.STRING,
    name: DataTypes.STRING,
    description: DataTypes.STRING,
    from_web_site: DataTypes.STRING,
    new_chapter: DataTypes.STRING,
    word_count: DataTypes.STRING,
    chapter_count: DataTypes.INTEGER,
    book_list_refer: DataTypes.INTEGER,
    tags: DataTypes.STRING,
    tag_ids: DataTypes.STRING
  }, {});
  book.associate = function(models) {
    // associations can be defined here
  };
  return book;
};