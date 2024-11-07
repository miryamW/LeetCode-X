func createQuestion(description string, level int) (*mongo.InsertOneResult, error) {
	question := Question{
		Description: description,
		Level:       level,
	}

	result, err := questionCollection.InsertOne(context.Background(), question)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func getQuestionByID(id string) (*Question, error) {
	questionID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var question Question
	err = questionCollection.FindOne(context.Background(), bson.M{"_id": questionID}).Decode(&question)
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func getAllQuestions() ([]Question, error) {
	cursor, err := questionCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var questions []Question
	for cursor.Next(context.Background()) {
		var question Question
		if err := cursor.Decode(&question); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return questions, nil
}

func updateQuestion(id string, description string, level int) (*mongo.UpdateResult, error) {
	questionID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"description": description,
			"level":       level,
		},
	}

	result, err := questionCollection.UpdateOne(context.Background(), bson.M{"_id": questionID}, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func deleteQuestion(id string) (*mongo.DeleteResult, error) {
	questionID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result, err := questionCollection.DeleteOne(context.Background(), bson.M{"_id": questionID})
	if err != nil {
		return nil, err
	}

	return result, nil
}



