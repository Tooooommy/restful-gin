'use strict';
module.exports = (sequelize, DataTypes) => {
  const book_list = sequelize.define('book_list', {
    name: DataTypes.STRING,
    description: DataTypes.STRING,
    owner_id: DataTypes.INTEGER,
    tags: DataTypes.STRING,
    tag_ids: DataTypes.STRING
  }, {});
  book_list.associate = function(models) {
    // associations can be defined here
  };
  return book_list;
};