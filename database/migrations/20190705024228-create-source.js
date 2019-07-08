'use strict';
const table = 'sources';
module.exports = {
  up: (queryInterface, Sequelize) => {
    return queryInterface.createTable(table, {
      id: {
        allowNull: false,
        autoIncrement: true,
        primaryKey: true,
        type: Sequelize.INTEGER
      },
      book_id: {
        allowNull: false,
        type: Sequelize.INTEGER,
        references: {
          model: 'books',
          key: 'id'
        }
      },
      name: {
        allowNull: false,
        type: Sequelize.STRING
      },
      url: {
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
      comment: '书籍来源',
      collate: 'utf8mb4_general_ci'
    })
        .then(()=>queryInterface.addIndex(table, ['deleted_at']));
  },
  down: (queryInterface, Sequelize) => {
    return queryInterface.removeIndex(table, ['deleted_at'])
        .then(()=>queryInterface.dropTable(table));
  }
};