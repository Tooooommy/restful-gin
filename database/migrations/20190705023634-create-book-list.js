'use strict';
const table = 'book_lists';
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
      description: {
        allowNull: false,
        type: Sequelize.STRING
      },
      owner_id: {
        allowNull: false,
        type: Sequelize.INTEGER,
        references: {
          model: 'accounts',
          key: 'id'
        }
      },
      tags: {
        allowNull: false,
        type: Sequelize.STRING
      },
      tag_ids: {
        allowNull: false,
        type: Sequelize.STRING
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
      comment: '书单列表',
      collate: 'utf8mb4_general_ci'
    })
        .then(()=>queryInterface.addIndex(table, ['name']))
        .then(()=>queryInterface.addIndex(table, ['deleted_at']))
  },
  down: (queryInterface, Sequelize) => {
    return queryInterface.removeIndex(table, ['name'])
        .then(()=>queryInterface.removeIndex(table, ['deleted_at']))
        .then(()=>queryInterface.dropTable('book_lists'));
  }
};