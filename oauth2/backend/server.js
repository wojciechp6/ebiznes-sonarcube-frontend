// server.js
const express = require('express');
const cors = require('cors');
const passport = require('passport');
const session = require('express-session');
const { sequelize, User } = require('./models');
const authRouter = require('./auth');
const config = require('./config');

const app = express();

app.use(cors({
  origin: process.env.FRONTEND_URL || 'http://localhost:3000',
  credentials: true,
}));
app.use(express.json());
app.use(session({ secret: config.JWT_SECRET, resave: false, saveUninitialized: false }));
app.use(passport.initialize());
app.use(passport.session());

passport.serializeUser((user, done) => done(null, user.id));
passport.deserializeUser(async (id, done) => {
  try {
    const user = await User.findByPk(id);
    done(null, user);
  } catch (err) {
    done(err);
  }
});

app.use('/auth', authRouter);

// Przykładowy endpoint chroniony
const jwt = require('jsonwebtoken');
app.get('/profile', async (req, res) => {
  const auth = req.headers.authorization;
  if (!auth) return res.status(401).json({ error: 'No token' });
  try {
    const decoded = jwt.verify(auth.split(' ')[1], config.JWT_SECRET);
    const user = await User.findByPk(decoded.id);
    if (!user) return res.status(404).json({ error: 'User not found' });
    res.json({ id: user.id, username: user.username, email: user.email, name: user.name, provider: user.provider });
  } catch (err) {
    res.status(401).json({ error: 'Invalid token' });
  }
});

// Synchronizacja bazy i start serwera
sequelize.sync().then(() => {
  app.listen(4000, () => {
    console.log('Server running on http://localhost:4000');
  });
});