// models.js
const { Sequelize, DataTypes } = require('sequelize');
const sequelize = new Sequelize({
  dialect: 'sqlite',
  storage: 'database.sqlite',
});

const User = sequelize.define('User', {
  id: { type: DataTypes.INTEGER, autoIncrement: true, primaryKey: true },
  username: { type: DataTypes.STRING, unique: true },
  password: { type: DataTypes.STRING }, // hash
  provider: { type: DataTypes.STRING }, // 'local', 'google', 'github'
  providerId: { type: DataTypes.STRING }, 
  email: { type: DataTypes.STRING },
  name: { type: DataTypes.STRING },
});

const Token = sequelize.define('Token', {
  token: { type: DataTypes.STRING, unique: true },
  userId: { type: DataTypes.INTEGER, allowNull: false },
});

User.hasMany(Token, { foreignKey: 'userId' });
Token.belongsTo(User, { foreignKey: 'userId' });

module.exports = { sequelize, User, Token };