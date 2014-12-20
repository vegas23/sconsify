package sconsify

import (
	"math/rand"
	"testing"
)

func TestNormalMode(t *testing.T) {
	playlists := InitPlaylists()
	if !playlists.isNormalMode() {
		t.Errorf("Playlists initial state should be Normal")
	}

	playlists.AddPlaylist("name", CreateDummyPlaylist())
	playlists.SetCurrents("name", 0)

	if track, repeating := playlists.GetNext(); track.Uri != "1" || repeating {
		t.Errorf("Random track should be 1")
	}
	if track, repeating := playlists.GetNext(); track.Uri != "2" || repeating {
		t.Errorf("Random track should be 2")
	}
	if track, repeating := playlists.GetNext(); track.Uri != "3" || repeating {
		t.Errorf("Random track should be 3")
	}
	if track, repeating := playlists.GetNext(); track.Uri != "0" || repeating {
		t.Errorf("Random track should be 0")
	}

	// normal mode doesn't support repeating flag
	if track, repeating := playlists.GetNext(); track.Uri != "1" || repeating {
		t.Errorf("Random track should be 1")
	}
}

func TestRandomMode(t *testing.T) {
	rand.Seed(123456789) // repeatable

	playlists := InitPlaylists()
	if !playlists.isNormalMode() {
		t.Errorf("Playlists initial state should be Normal")
	}

	playlists.AddPlaylist("name", CreateDummyPlaylist())
	playlists.SetCurrents("name", 0)
	playlists.SetMode(RandomMode)

	order := []string{"3", "0", "2", "1"}
	for _, expectedUri := range order {
		if track, repeating := playlists.GetNext(); expectedUri != track.Uri || repeating {
			t.Errorf("Random track should be %v and not repeating but it is %v and isRepeating? %v", track.Uri, repeating)
		}
	}

	// now is repeating
	if track, repeating := playlists.GetNext(); track.Uri != "3" || !repeating {
		t.Errorf("Random track should be 3 and repeating but it is %v and isRepeating? %v", repeating)
	}
}

func TestAllRandomMode(t *testing.T) {
	rand.Seed(123456789) // repeatable

	playlists := InitPlaylists()
	if !playlists.isNormalMode() {
		t.Errorf("Playlists initial state should be Normal")
	}

	playlists.AddPlaylist("name", CreateDummyPlaylist())
	playlists.AddPlaylist("name1", CreateDummyPlaylist())
	playlists.SetCurrents("name", 0)
	playlists.SetMode(AllRandomMode)

	order := []string{"3", "3", "2", "1", "0", "1", "2", "0"}

	for _, expectedUri := range order {
		if track, repeating := playlists.GetNext(); expectedUri != track.Uri || repeating {
			t.Errorf("Random track should be %v and not repeating but it is %v and isRepeating? %v", expectedUri, track.Uri, repeating)
		}
	}

	// now is repeating
	if track, repeating := playlists.GetNext(); track.Uri != "3" || !repeating {
		t.Errorf("Random track should be 3 and repeating but it is %v and isRepeating? %v", track.Uri, repeating)
	}
}

func TestSequentialRandomMode(t *testing.T) {
	rand.Seed(123456789) // repeatable

	playlists := InitPlaylists()
	if !playlists.isNormalMode() {
		t.Errorf("Playlists initial state should be Normal")
	}

	playlists.AddPlaylist("name", CreateDummyPlaylist())
	playlists.AddPlaylist("name1", CreateDummyPlaylist())
	playlists.SetCurrents("name", 0)
	playlists.SetMode(SequentialMode)

	order := []string{"0", "1", "2", "3", "0", "1", "2", "3"}

	for _, expectedUri := range order {
		if track, repeating := playlists.GetNext(); expectedUri != track.Uri || repeating {
			t.Errorf("Random track should be %v and not repeating but it is %v and isRepeating? %v", expectedUri, track.Uri, repeating)
		}
	}

	// now is repeating
	if track, repeating := playlists.GetNext(); track.Uri != "0" || !repeating {
		t.Errorf("Random track should be 0 and repeating but it is %v and isRepeating? %v", track.Uri, repeating)
	}
}

func TestPremadeTracks(t *testing.T) {
	playlists := InitPlaylists()
	if playlists.PremadeTracks() != 0 {
		t.Errorf("PremadeTracks should be empty")
	}

	playlists.SetMode(SequentialMode)
	if playlists.PremadeTracks() != 0 {
		t.Errorf("PremadeTracks should be empty")
	}

	playlists = InitPlaylists()
	playlists.AddPlaylist("name", CreateDummyPlaylist())
	playlists.SetMode(SequentialMode)

	if playlists.PremadeTracks() != 4 {
		t.Errorf("PremadeTracks should be 4")
	}

	playlists.AddPlaylist("name1", CreateDummyPlaylist())
	if playlists.PremadeTracks() != 8 {
		t.Errorf("PremadeTracks should be 8")
	}
}

func TestSetCurrents(t *testing.T) {
	playlists := InitPlaylists()

	if err := playlists.SetCurrents("not to be found", 10); err == nil {
		t.Errorf("Playlist should not be found")
	}
	if playlists.HasPlaylistSelected() {
		t.Errorf("No playlist should be selected")
	}

	playlists.AddPlaylist("name", CreateDummyPlaylist())

	if err := playlists.SetCurrents("name", 0); err != nil {
		t.Errorf("Playlist and track should be found")
	}
	if err := playlists.SetCurrents("name", 3); err != nil {
		t.Errorf("Playlist and track should be found")
	}
	if !playlists.HasPlaylistSelected() {
		t.Errorf("It has playlist selected")
	}
}

func TestTracks(t *testing.T) {
	playlists := InitPlaylists()

	if playlists.Tracks() != 0 {
		t.Errorf("Tracks should be empty")
	}

	playlists.AddPlaylist("name", CreateDummyPlaylist())
	if playlists.Tracks() != 4 {
		t.Errorf("Tracks should be 4")
	}

	playlists.AddPlaylist("name1", CreateDummyPlaylist())
	if playlists.Tracks() != 8 {
		t.Errorf("Tracks should be 8")
	}
}

func TestGetNames(t *testing.T) {
	playlists := InitPlaylists()

	if len(playlists.GetNames()) != 0 {
		t.Errorf("Playlists should be empty")
	}

	playlists.AddPlaylist("name", CreateDummyPlaylist())
	names := playlists.GetNames()
	if len(names) != 1 {
		t.Errorf("Should have only one name")
	}

	playlists.AddPlaylist("name1", CreateDummyPlaylist())
	names = playlists.GetNames()
	if len(names) != 2 {
		t.Errorf("Should have 2 names")
	}
}

func TestGetNext(t *testing.T) {
	playlists := InitPlaylists()

	if track, _ := playlists.GetNext(); track != nil {
		t.Errorf("Track should not be found")
	}

	playlists.AddPlaylist("name", CreateDummyPlaylist())

	playlists.SetCurrents("name", 0)

	if track, _ := playlists.GetNext(); track != nil && track.Uri != "1" {
		t.Errorf("Next track should be 1")
	}
	if track, _ := playlists.GetNext(); track != nil && track.Uri != "2" {
		t.Errorf("Next track should be 2")
	}
	if track, _ := playlists.GetNext(); track != nil && track.Uri != "3" {
		t.Errorf("Next track should be 3")
	}
}

func TestPlaylists(t *testing.T) {
	playlists := InitPlaylists()

	if playlists.Playlists() != 0 {
		t.Errorf("Playlist count should be empty")
	}

	playlists.AddPlaylist("name", CreateDummyPlaylist())

	if count := playlists.Playlists(); count != 1 {
		t.Errorf("Playlist count should be 1 but it is %v", count)
	}

	playlists.AddPlaylist("name1", CreateDummyPlaylist())
	if count := playlists.Playlists(); count != 2 {
		t.Errorf("Playlist count should be 2 but it is %v", count)
	}
}

func TestStateMode(t *testing.T) {
	playlists := InitPlaylists()

	if mode := playlists.GetModeAsString(); mode != "" {
		t.Errorf("Empty playlists should be on empty state but it is [%v]", mode)
	}

	playlists.playMode = RandomMode
	if mode := playlists.GetModeAsString(); mode != "[Random] " {
		t.Errorf("Mode was set to true but is returning %v", mode)
	}

	playlists.playMode = AllRandomMode
	if mode := playlists.GetModeAsString(); mode != "[All Random] " {
		t.Errorf("Mode was set to true but is returning %v", mode)
	}

	playlists.playMode = NormalMode
	if mode := playlists.GetModeAsString(); mode != "" {
		t.Errorf("Mode was set to false but is returning %v", mode)
	}
}

func TestStateInvertMode(t *testing.T) {
	playlists := InitPlaylists()

	playlists.InvertMode(RandomMode)
	if mode := playlists.GetModeAsString(); mode != "[Random] " {
		t.Errorf("Mode was inverted to random but is returning %v", mode)
	}

	playlists.InvertMode(RandomMode)
	if mode := playlists.GetModeAsString(); mode != "" {
		t.Errorf("Mode was inverted to normal but is returning %v", mode)
	}

	playlists.InvertMode(AllRandomMode)
	if mode := playlists.GetModeAsString(); mode != "[All Random] " {
		t.Errorf("Mode was inverted to allRandom but is returning %v", mode)
	}

	playlists.InvertMode(AllRandomMode)
	if mode := playlists.GetModeAsString(); mode != "" {
		t.Errorf("Mode was inverted to random but is returning %v", mode)
	}
}