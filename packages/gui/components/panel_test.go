package components

import (
	"testing"
)

func TestBasePanel_ScrollingBehavior(t *testing.T) {
	config := BasePanelConfig{
		Name:        "test",
		Title:       "Test Panel",
		Highlighted: true,
		Selectable:  true,
	}
	
	panel := NewBasePanel(config)
	
	testItems := []string{
		"item1", "item2", "item3", "item4", "item5",
		"item6", "item7", "item8", "item9", "item10",
		"item11", "item12", "item13", "item14", "item15",
	}
	panel.Items = testItems
	
	if panel.SelectedIndex != 0 {
		t.Errorf("Expected initial selected index to be 0, got %d", panel.SelectedIndex)
	}
	if panel.ScrollOffset != 0 {
		t.Errorf("Expected initial scroll offset to be 0, got %d", panel.ScrollOffset)
	}
	
	viewHeight := 5
	
	for i := 0; i < 7; i++ {
		panel.MoveDownWithScroll(viewHeight)
	}
	
	if panel.SelectedIndex != 7 {
		t.Errorf("Expected selected index to be 7, got %d", panel.SelectedIndex)
	}
	
	if panel.ScrollOffset != 3 {
		t.Errorf("Expected scroll offset to be 3, got %d", panel.ScrollOffset)
	}
	
	for i := 0; i < 3; i++ {
		panel.MoveUpWithScroll(viewHeight)
	}
	
	if panel.SelectedIndex != 4 {
		t.Errorf("Expected selected index to be 4, got %d", panel.SelectedIndex)
	}
	
	visibleItems := panel.GetVisibleItems(viewHeight)
	if len(visibleItems) != viewHeight {
		t.Errorf("Expected %d visible items, got %d", viewHeight, len(visibleItems))
	}
	
	visibleSelectedIndex := panel.GetVisibleSelectedIndex()
	if visibleSelectedIndex < 0 || visibleSelectedIndex >= len(visibleItems) {
		t.Errorf("Visible selected index %d is out of bounds for visible items", visibleSelectedIndex)
	}
}

func TestBasePanel_ScrollingEdgeCases(t *testing.T) {
	config := BasePanelConfig{
		Name:        "test",
		Title:       "Test Panel",
		Highlighted: true,
		Selectable:  true,
	}
	
	panel := NewBasePanel(config)
	
	panel.Items = []string{}
	panel.MoveDownWithScroll(5)
	if panel.SelectedIndex != 0 {
		t.Errorf("Expected selected index to remain 0 with empty items, got %d", panel.SelectedIndex)
	}
	
	panel.Items = []string{"single"}
	panel.MoveDownWithScroll(5)
	if panel.SelectedIndex != 0 {
		t.Errorf("Expected selected index to remain 0 with single item, got %d", panel.SelectedIndex)
	}
	
	panel.Items = []string{"item1", "item2", "item3"}
	panel.MoveDownWithScroll(5)
	panel.MoveDownWithScroll(5)
	panel.MoveDownWithScroll(5)
	if panel.SelectedIndex != 2 {
		t.Errorf("Expected selected index to be 2 (last item), got %d", panel.SelectedIndex)
	}
	if panel.ScrollOffset != 0 {
		t.Errorf("Expected scroll offset to be 0 when items fit in view, got %d", panel.ScrollOffset)
	}
}