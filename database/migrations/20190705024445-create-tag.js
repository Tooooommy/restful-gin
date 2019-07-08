'use strict';
const table = 'tags';
module.exports = {
  up: (queryInterface, Sequelize) => {
    return queryInterface.createTable(table, {
      id: {
        allowNull: false,
        autoIncrement: true,
        primaryKey: true,
        type: Sequelize.INTEGER
      },
      name: {
        allowNull: false,
        type: Sequelize.STRING
      },
      type: {
        allowNull: false,
        type: Sequelize.INTEGER,
        defaultValue: 0
      },
      book_list_refer: {
        allowNull: false,
        type: Sequelize.INTEGER,
        defaultValue: 0
      },
      book_refer: {
        allowNull: false,
        type: Sequelize.INTEGER,
        defaultValue: 0
      },
      created_at: {
        allowNull: false,
        type: Sequelize.DATE
      },
      updated_at: {
        allowNull: false,
        type: Sequelize.DATE
      },
      deleted_at: {
        type: Sequelize.DATE
      }
    }, {
      engine: 'InnoDB',
      charset: 'utf8mb4',
      comment: '标签',
      collate: 'utf8mb4_general_ci'
    })
        .then(()=>queryInterface.addIndex(table, ['name']))
        .then(()=>queryInterface.addIndex(table, ['deleted_at']));
  },
  down: (queryInterface, Sequelize) => {
    return queryInterface.removeIndex(table, ['name'])
        .then(()=>queryInterface.removeIndex(table, ['deleted_at']))
        .then(()=>queryInterface.dropTable('tags'));
  }
};