package shaerlock

import "testing"

func TestRegisterSleuth(t *testing.T) {
	initialLen := len(sleuths)
	err := RegisterSleuth("torrent", func(path string) (Media, error){
		return Media{}, nil
	})
	if err != nil {
		t.Errorf("Registering sleuth for torrent files shouldn't fail: %#v", err)
	}
	if initialLen + 1 != len(sleuths) {
		t.Error("Registering sleuth for torrent files should increment sleuth registry size by one")
	}
	initialLen = len(sleuths)
	err = RegisterSleuth("torrent", func(path string) (Media, error){
		return Media{}, nil
	})
	if err == nil {
		t.Error("Registering sleuth for torrent files should trigger a duplicate error")
	}
	if initialLen != len(sleuths) {
		t.Error("Registering sleuth for torrent files should not modify the length of the sleuth registry")
	}
}