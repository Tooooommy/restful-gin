'use strict';
module.exports = (sequelize, DataTypes) => {
  const account = sequelize.define('account', {
    avatar: DataTypes.STRING,
    name: DataTypes.STRING,
    email: DataTypes.STRING,
    password: DataTypes.STRING,
    session_id: DataTypes.STRING,
    status: DataTypes.BOOLEAN
  }, {});
  account.associate = function(models) {
    // associations can be defined here
  };
  return account;
};