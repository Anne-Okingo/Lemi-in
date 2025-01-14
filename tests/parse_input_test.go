package test

import (
	"os"
	"testing"

	"lemin/models"
	"lemin/utils"
)

func TestParseInput(t *testing.T) {
	// Create a temporary file with valid input
	content := `3
##start
0 1 0
##end
1 5 0
2 9 0
0-2
2-1
`
	tmpfile, err := os.CreateTemp("", "test_input_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}
	// Call parseInput with the temporary file
	ants, rooms, links, err := utils.ParseInput(tmpfile.Name())
	// Check for errors
	if err != nil {
		t.Errorf("parseInput returned an error: %v", err)
	}
	// Check the number of ants
	if ants != 3 {
		t.Errorf("Expected 3 ants, got %d", ants)
	}
	// Check the number of rooms
	expectedRooms := 3
	if len(rooms) != expectedRooms {
		t.Errorf("Expected %d rooms, got %d", expectedRooms, len(rooms))
	}
	// Check the number of links
	expectedLinks := 2
	if len(links) != expectedLinks {
		t.Errorf("Expected %d links, got %d", expectedLinks, len(links))
	}
	// Check if start and end rooms are correctly marked
	startRoom := rooms[0]
	if !startRoom.IsStart || startRoom.Name != "0" {
		t.Errorf("Start room not correctly identified")
	}
	endRoom := rooms[1]
	if !endRoom.IsEnd || endRoom.Name != "1" {
		t.Errorf("End room not correctly identified")
	}
	// Check if links are correctly parsed
	expectedLinkPairs := [][2]string{{"0", "2"}, {"2", "1"}}
	for i, link := range links {
		if link.From != expectedLinkPairs[i][0] || link.To != expectedLinkPairs[i][1] {
			t.Errorf("Link %d not correctly parsed. Expected %v-%v, got %v-%v",
				i, expectedLinkPairs[i][0], expectedLinkPairs[i][1], link.From, link.To)
		}
	}
}

func TestParseInput_EmptyFile(t *testing.T) {
	// Create a temporary empty file
	tmpfile, err := os.CreateTemp("", "empty_test_file")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	// Close the file immediately as we want it to be empty
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temporary file: %v", err)
	}
	// Call parseInput with the empty file
	ants, rooms, links, err := utils.ParseInput(tmpfile.Name())
	// Check the results
	if err == nil {
		t.Errorf("Expected error for empty file, got: %v", nil)
	}
	if ants != 0 {
		t.Errorf("Expected 0 ants for empty file, got: %d", ants)
	}
	if len(rooms) != 0 {
		t.Errorf("Expected 0 rooms for empty file, got: %d", len(rooms))
	}
	if len(links) != 0 {
		t.Errorf("Expected 0 links for empty file, got: %d", len(links))
	}
}

func TestParseInput_InvalidAnts(t *testing.T) {
	// Create a temporary file with invalid input
	tmpfile, err := os.CreateTemp("", "test_invalid_ants")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	// Write invalid input to the file
	_, err = tmpfile.WriteString("0\n")
	if err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()
	// Call parseInput with the temporary file
	ants, rooms, links, err := utils.ParseInput(tmpfile.Name())
	// Check the results
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if ants != 0 || rooms != nil || links != nil {
		t.Errorf("Expected (0, nil, nil) for invalid input, but got (%d, %v, %v)", ants, rooms, links)
	}
	if err.Error() != "ERROR: invalid data format, invalid number of Ants" {
		t.Errorf("Expected error message 'ERROR: invalid data format, invalid number of Ants', but got '%s'", err.Error())
	}
}

func TestParseInput_IdentifiesStartRoom(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test_input_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	input := `3
##start
start 0 1
room1 2 3
##end
room2 4 5
start-room1
start-room2
`
	if _, err := tempFile.WriteString(input); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tempFile.Close()
	ants, rooms, _, err := utils.ParseInput(tempFile.Name())
	if err != nil {
		t.Fatalf("parseInput returned an error: %v", err)
	}
	if ants != 3 {
		t.Errorf("Expected 3 ants, got %d", ants)
	}
	startRoom := findStartRoom(rooms)
	if startRoom == nil {
		t.Fatalf("No start room found")
	}
	if startRoom.Name != "start" || startRoom.X != 0 || startRoom.Y != 1 || !startRoom.IsStart {
		t.Errorf("Start room not correctly identified: %+v", startRoom)
	}
}

func findStartRoom(rooms []models.Room) *models.Room {
	for _, room := range rooms {
		if room.IsStart {
			return &room
		}
	}
	return nil
}

func TestParse_InputEndRoom(t *testing.T) {
	// Create a temporary file with test input
	tmpfile, err := os.CreateTemp("", "test_input_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	// Write test data to the temporary file
	testInput := `3
##start
start 0 1
middle 1 1
##end
end 2 1
start-middle
middle-end
`
	if _, err := tmpfile.Write([]byte(testInput)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}
	// Call parseInput with the temporary file
	ants, rooms, _, err := utils.ParseInput(tmpfile.Name())
	// Check for errors
	if err != nil {
		t.Fatalf("parseInput returned an error: %v", err)
	}
	// Check if the number of ants is correct
	if ants != 3 {
		t.Errorf("Expected 3 ants, got %d", ants)
	}
	// Check if the end room is correctly marked
	endRoom := rooms[2] // Assuming the end room is the last one added
	if !endRoom.IsEnd {
		t.Errorf("End room not correctly marked: %+v", endRoom)
	}
	// Check if the end room has the correct name
	if endRoom.Name != "end" {
		t.Errorf("Expected end room name to be 'end', got '%s'", endRoom.Name)
	}
}

func TestParseInput_IgnoresCommentsAndEmptyLines(t *testing.T) {
	// Create a temporary file with test input
	tmpfile, err := os.CreateTemp("", "test_input_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	// Write test input to the temporary file
	testInput := `5
# This is a comment
##start
room1 0 1
##end
room2 2 0
# Another comment
room3 1 1
room1-room2
room1-room3
`
	if _, err := tmpfile.Write([]byte(testInput)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}
	// Call parseInput with the temporary file
	ants, rooms, links, err := utils.ParseInput(tmpfile.Name())
	// Check for errors
	if err != nil {
		t.Fatalf("parseInput returned an error: %v", err)
	}
	// Check the number of ants
	if ants != 5 {
		t.Errorf("Expected 5 ants, got %d", ants)
	}
	// Check the number of rooms
	expectedRooms := 3
	if len(rooms) != expectedRooms {
		t.Errorf("Expected %d rooms, got %d", expectedRooms, len(rooms))
	}
	// Check the number of links
	expectedLinks := 2
	if len(links) != expectedLinks {
		t.Errorf("Expected %d links, got %d", expectedLinks, len(links))
	}
	// Check if start and end rooms are correctly marked
	for _, room := range rooms {
		if room.Name == "room1" && !room.IsStart {
			t.Errorf("Expected room1 to be marked as start")
		}
		if room.Name == "room2" && !room.IsEnd {
			t.Errorf("Expected room2 to be marked as end")
		}
	}
}
