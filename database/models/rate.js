'use strict';
module.exports = (sequelize, DataTypes) => {
  const rate = sequelize.define('rate', {
    book_id: DataTypes.INTEGER,
    account_id: DataTypes.INTEGER,
    rate: DataTypes.INTEGER
  }, {});
  rate.associate = function(models) {
    // associations can be defined here
  };
  return rate;
};