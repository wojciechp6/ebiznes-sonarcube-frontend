// auth.js
const express = require('express');
const passport = require('passport');
const LocalStrategy = require('passport-local').Strategy;
const GoogleStrategy = require('passport-google-oauth20').Strategy;
const GitHubStrategy = require('passport-github2').Strategy;
const jwt = require('jsonwebtoken');
const bcrypt = require('bcryptjs');
const { User, Token } = require('./models');
const config = require('./config');

const router = express.Router();

// Passport local strategy
passport.use(new LocalStrategy(async (username, password, done) => {
  try {
    const user = await User.findOne({ where: { username, provider: 'local' } });
    if (!user) return done(null, false, { message: 'Incorrect username.' });
    const valid = await bcrypt.compare(password, user.password);
    if (!valid) return done(null, false, { message: 'Incorrect password.' });
    return done(null, user);
  } catch (err) {
    return done(err);
  }
}));

// Passport Google strategy
passport.use(new GoogleStrategy({
  clientID: config.GOOGLE_CLIENT_ID,
  clientSecret: config.GOOGLE_CLIENT_SECRET,
  callbackURL: config.BASE_URL + '/auth/google/callback',
}, async (accessToken, refreshToken, profile, done) => {
  try {
    let user = await User.findOne({ where: { provider: 'google', providerId: profile.id } });
    if (!user) {
      user = await User.create({
        provider: 'google',
        providerId: profile.id,
        email: profile.emails[0].value,
        name: profile.displayName,
      });
    }
    return done(null, user);
  } catch (err) {
    return done(err);
  }
}));

// Passport Github strategy
passport.use(new GitHubStrategy({
  clientID: config.GITHUB_CLIENT_ID,
  clientSecret: config.GITHUB_CLIENT_SECRET,
  callbackURL: config.BASE_URL + '/auth/github/callback',
}, async (accessToken, refreshToken, profile, done) => {
  try {
    let user = await User.findOne({ where: { provider: 'github', providerId: profile.id } });
    if (!user) {
      user = await User.create({
        provider: 'github',
        providerId: profile.id,
        email: profile.emails ? profile.emails[0].value : null,
        name: profile.displayName,
      });
    }
    return done(null, user);
  } catch (err) {
    return done(err);
  }
}));

// JWT token generator
async function issueToken(user) {
  const token = jwt.sign({ id: user.id, username: user.username, provider: user.provider }, config.JWT_SECRET, { expiresIn: '1h' });
  await Token.create({ token, userId: user.id });
  return token;
}

// Rejestracja
router.post('/register', async (req, res) => {
  const { username, password, email, name } = req.body;
  if (!username || !password) return res.status(400).json({ error: 'Missing fields' });
  const hash = await bcrypt.hash(password, 10);
  try {
    const user = await User.create({ username, password: hash, provider: 'local', email, name });
    res.json({ success: true });
  } catch (err) {
    res.status(400).json({ error: 'User exists' });
  }
});

// Logowanie klasyczne
router.post('/login', (req, res, next) => {
  passport.authenticate('local', async (err, user, info) => {
    if (err) return next(err);
    if (!user) return res.status(401).json({ error: info.message });
    const token = await issueToken(user);
    res.json({ token });
  })(req, res, next);
});

// Google OAuth2
router.get('/google', passport.authenticate('google', { scope: ['profile', 'email'] }));
router.get('/google/callback', passport.authenticate('google', { session: false, failureRedirect: '/' }), async (req, res) => {
  const token = await issueToken(req.user);
  // Przekierowanie do frontendu z tokenem w query
  res.redirect(`${process.env.FRONTEND_URL || 'http://localhost:3000'}/oauth2?token=${token}`);
});

// Github OAuth2
router.get('/github', passport.authenticate('github', { scope: ['user:email'] }));
router.get('/github/callback', passport.authenticate('github', { session: false, failureRedirect: '/' }), async (req, res) => {
  const token = await issueToken(req.user);
  res.redirect(`${process.env.FRONTEND_URL || 'http://localhost:3000'}/oauth2?token=${token}`);
});

module.exports = router;