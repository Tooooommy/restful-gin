'use strict';
module.exports = (sequelize, DataTypes) => {
  const user_follow_book_list = sequelize.define('user_follow_book_list', {
    account_id: DataTypes.INTEGER,
    book_list_id: DataTypes.INTEGER
  }, {});
  user_follow_book_list.associate = function(models) {
    // associations can be defined here
  };
  return user_follow_book_list;
};