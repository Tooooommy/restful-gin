'use strict';
module.exports = (sequelize, DataTypes) => {
  const comment = sequelize.define('comment', {
    book_list_id: DataTypes.INTEGER,
    account_id: DataTypes.INTEGER,
    content: DataTypes.STRING,
    reply_id: DataTypes.INTEGER
  }, {});
  comment.associate = function(models) {
    // associations can be defined here
  };
  return comment;
};