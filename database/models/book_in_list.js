'use strict';
module.exports = (sequelize, DataTypes) => {
  const book_in_list = sequelize.define('book_in_list', {
    book_id: DataTypes.INTEGER,
    book_list_id: DataTypes.INTEGER
  }, {});
  book_in_list.associate = function(models) {
    // associations can be defined here
  };
  return book_in_list;
};