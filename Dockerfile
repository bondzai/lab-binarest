# Use the official Node.js 14 image as the base
FROM node:14

# Install Docker inside the container
RUN apt-get update && apt-get install -y docker.io

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package.json package-lock.json ./

# Install Node.js dependencies
RUN npm install --production

# Copy the rest of the application code
COPY . .

# Start the Node.js application
CMD ["npm", "start"]
