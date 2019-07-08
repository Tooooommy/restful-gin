'use strict';
module.exports = (sequelize, DataTypes) => {
  const tag = sequelize.define('tag', {
    name: DataTypes.STRING,
    type: DataTypes.INTEGER,
    book_list_refer: DataTypes.INTEGER,
    book_refer: DataTypes.INTEGER
  }, {});
  tag.associate = function(models) {
    // associations can be defined here
  };
  return tag;
};