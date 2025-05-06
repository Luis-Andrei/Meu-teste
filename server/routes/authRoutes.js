const express = require('express');
const router = express.Router();
const { signup, login, getMe } = require('../controllers/authController');
const validate = require('../middleware/validationMiddleware');
const { signupSchema, loginSchema } = require('../validations/authValidation');
const authMiddleware = require('../middleware/authMiddleware');

// Public routes
router.post('/signup', validate(signupSchema), signup);
router.post('/login', validate(loginSchema), login);

// Protected routes
router.get('/me', authMiddleware, getMe);

module.exports = router; 