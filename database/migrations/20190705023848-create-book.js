'use strict';
const table = 'books';
module.exports = {
  up: (queryInterface, Sequelize) => {
    return queryInterface.createTable(table, {
      id: {
        allowNull: false,
        autoIncrement: true,
        primaryKey: true,
        type: Sequelize.INTEGER
      },
      cover: {
        allowNull: false,
        type: Sequelize.STRING
      },
      author: {
        allowNull: false,
        type: Sequelize.STRING
      },
      name: {
        allowNull: false,
        type: Sequelize.STRING
      },
      description: {
        allowNull: false,
        type: Sequelize.STRING
      },
      from_web_site: {
        allowNull: false,
        type: Sequelize.STRING
      },
      new_chapter: {
        allowNull: false,
        type: Sequelize.STRING
      },
      word_count: {
        allowNull: false,
        defaultValue: '0',
        type: Sequelize.STRING
      },
      chapter_count: {
        allowNull: false,
        defaultValue: 0,
        type: Sequelize.INTEGER
      },
      book_list_refer: {
        allowNull: false,
        defaultValue: 0,
        type: Sequelize.INTEGER
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
      comment: '书籍',
      collate: 'utf8mb4_general_ci'
    })
        .then(()=>queryInterface.addIndex(table, ['name']))
        .then(()=>queryInterface.addIndex(table, ['author']))
        .then(()=>queryInterface.addIndex(table, ['deleted_at']))
  },
  down: (queryInterface, Sequelize) => {
    return queryInterface.removeIndex(table, ['name'])
        .then(()=>queryInterface.removeIndex(table, ['author']))
        .then(()=>queryInterface.removeIndex(table, ['deleted_at']))
        .then(()=>queryInterface.dropTable(table));
  }
};