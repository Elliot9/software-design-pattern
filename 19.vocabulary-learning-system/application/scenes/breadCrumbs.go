package scenes

import (
	"fmt"
	"github/elliot9/class19/pkg/sceneManagement"
	"strings"
)

type Breadcrumb interface {
	Render() []string
}

type SimpleBreadcrumb struct {
	path             []string
	navigationPrefix string
}

func NewSimpleBreadcrumb(path []string) *SimpleBreadcrumb {
	return &SimpleBreadcrumb{path: path, navigationPrefix: "/"}
}

func (b *SimpleBreadcrumb) Render() []string {
	return []string{fmt.Sprintf("%s %s", b.navigationPrefix, strings.Join(b.path, fmt.Sprintf(" %s ", b.navigationPrefix)))}
}

type NavigationBreadcrumb struct {
	sceneManager     *sceneManagement.SceneManager
	navigationPrefix string
}

func NewNavigationBreadcrumb(sceneManager *sceneManagement.SceneManager) *NavigationBreadcrumb {
	return &NavigationBreadcrumb{
		sceneManager:     sceneManager,
		navigationPrefix: "/",
	}
}

func (b *NavigationBreadcrumb) Render() []string {
	fullPath := b.sceneManager.GetFullPath(b.sceneManager.Current().GetName())

	// 根目錄不顯示
	if len(fullPath) <= 1 {
		return []string{b.navigationPrefix}
	}

	return []string{fmt.Sprintf("%s %s", b.navigationPrefix, strings.Join(fullPath[1:], fmt.Sprintf(" %s ", b.navigationPrefix)))}
}
