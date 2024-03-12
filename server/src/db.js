import mongoose from 'mongoose';

const dbURI = 'mongodb://localhost:27017/sales-dashboard';

mongoose
  .connect(dbURI, {})
  .then(() => console.log('Connected to MongoDB'))
  .catch((error) => console.error('MongoDB connection error:', error));
