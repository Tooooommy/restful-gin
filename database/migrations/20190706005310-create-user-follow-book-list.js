'use strict';
const table = 'user_follow_book_lists';
module.exports = {
  up: (queryInterface, Sequelize) => {
    return queryInterface.createTable(table, {
      id: {
        allowNull: false,
        autoIncrement: true,
        primaryKey: true,
        type: Sequelize.INTEGER
      },
      account_id: {
        allowNull: false,
        type: Sequelize.INTEGER,
        references: {
          model: 'accounts',
          key: 'id'
        }
      },
      book_list_id: {
        allowNull: false,
        type: Sequelize.INTEGER,
        references: {
          model: 'book_lists',
          key: 'id'
        }
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
    },{
      engine: 'InnoDB',
      charset: 'utf8mb4',
      comment: '标签',
      collate: 'utf8mb4_general_ci'
  })
        .then(()=>queryInterface.addIndex(table, ['deleted_at']))
  },
  down: (queryInterface, Sequelize) => {
    return queryInterface.removeIndex(table, ['deleted_at'])
        .then(()=>queryInterface.dropTable(table));
  }
};