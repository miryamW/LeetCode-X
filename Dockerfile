# Use the basic Python image
FROM python:3.9

# Install pip, pytest, and any other required dependencies
RUN pip install --upgrade pip && pip install pytest

# Install Node.js and Jest if needed (if you have JavaScript tests)
RUN curl -sL https://deb.nodesource.com/setup_16.x | bash - \
    && apt-get install -y nodejs \
    && npm install -g jest

# Set the working directory inside the container
WORKDIR /app

# Copy the project files into the container
COPY . /app

# Default command to run pytest with a specific test file
CMD ["sh", "-c", "pytest tests/${TEST_ID}.py"]
