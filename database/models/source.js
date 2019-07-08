'use strict';
module.exports = (sequelize, DataTypes) => {
  const source = sequelize.define('source', {
    book_id: DataTypes.INTEGER,
    name: DataTypes.STRING,
    url: DataTypes.STRING
  }, {});
  source.associate = function(models) {
    // associations can be defined here
  };
  return source;
};