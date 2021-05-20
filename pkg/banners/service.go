package banners

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// Service npenctasnset co6oi cepsuc no ynpasnenwo OaHHepamn.
type Service struct {
	nextAccountID int64
	mu            sync.RWMutex

	items []*Banner
}

// NewService co3qaét cepsuc.
func NewService() *Service {
	return &Service{items: make([]*Banner, 0)}

}

// Banner npenctasnaet codoi GaHHep.
type Banner struct {
	ID int64

	Title string

	Content string

	Button string

	Link string

	Image string
}

// ByID Bo3BpawaeT OaHHep no upeHTHOuKaTopy.
func (s *Service) ByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, banner := range s.items {
		if banner.ID == id {
			return banner, nil
		}
	}

	return nil, errors.New("item not found")
}

// All for
func (s *Service) All(ctx context.Context) ([]*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	// for _, banner := range s.items {
	// 	if banner.ID == id {
	// 		return banner, nil
	// 	}
	// }
	// banners := s.items
	// if len(s.items) == 0 {
	// 	return nil, errors.New("no items found")
	// }
	return s.items, nil
	//panic("not implemented")
}

// Save for
func (s *Service) Save(ctx context.Context, item *Banner, file multipart.File) (*Banner, error) {
	//var lastID int64
	s.mu.RLock()
	defer s.mu.RUnlock()
	//	for _, banner := range s.items {
	if item.ID == 0 {
		// lenBanners := len(s.items) - 1
		// for i, banner := range s.items {
		// 	if i == lenBanners {
		// 		lastID = banner.ID
		// 	}
		// }
		s.nextAccountID++
		item.ID = s.nextAccountID
		//		item.ID = int64(len(s.items)) + 1
		nameImage := item.Image
		if nameImage != "" {
			extenIndex := strings.Index(nameImage, ".")
			fileExtension := nameImage[extenIndex:]
			item.Image = strconv.FormatInt(item.ID, 10) + fileExtension
		}
		saveFile(file, item)
		s.items = append(s.items, item)

		return item, nil
	}
	if item.ID != 0 {
		for _, banner := range s.items {
			if banner.ID == item.ID {
				banner.Button = item.Button
				banner.Content = item.Content
				banner.Link = item.Link
				banner.Title = item.Title
				if item.Image != "" {
					
					banner.Image = item.Image
					saveFile(file, item)
				}
				if item.Image == "" {
					item.Image = banner.Image
					//banner.Image = ""
					// nameImage := item.Image
					// extenIndex := strings.Index(nameImage, ".")
					// fileExtension := nameImage[extenIndex:]
					// item.Image = strconv.FormatInt(item.ID, 10) + fileExtension
					// banner.Image = item.Image
				}
				
				return item, nil
			}
		}
	}
	return nil, errors.New("item not found")

}

// RemoveByID for
func (s *Service) RemoveByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for i, banner := range s.items {
		if banner.ID == id {

			s.items = append(s.items[:i], s.items[i+1:]...)

			return banner, nil
		}
	}

	return nil, errors.New("item not found")
}

// Initial for
func (s *Service) Initial(request *http.Request) Banner {

	idParam := request.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Print(err)

	}

	titleParam := request.URL.Query().Get("title")
	contentParam := request.URL.Query().Get("content")
	buttonParam := request.URL.Query().Get("button")
	linkParam := request.URL.Query().Get("link")

	banner := Banner{
		ID: id,

		Title: titleParam,

		Content: contentParam,

		Button: buttonParam,

		Link: linkParam,

		Image: "image1",
	}

	// banner2 := Banner{
	// 	ID: 2,

	// 	Title: "Title New",

	// 	Content: "Content New",

	// 	Button: "Button New",

	// 	Link: "Link New",
	// }

	//item := s.items
	//	s.items = append(s.items, &banner)
	//s.items = append(s.items, &banner2)
	//item[1] = &banner
	//	panic("not implemented")

	return banner
}

func saveFile(fileA multipart.File, item *Banner) {
	content := make([]byte, 0)
	buf := make([]byte, 4)
		for {
			read, err := fileA.Read(buf)
			if err == io.EOF {
				break
			}
			content = append(content, buf[:read]...)
		}

		fileNameNew := item.Image
		if fileNameNew != "" {
			wdd1 := "web/banners" + "/" + fileNameNew
			//wdd1 := "c:/projects/http/web/banners" + "/" + fileNameNew
			//log.Print(wdd)
			err := ioutil.WriteFile(wdd1, content, 0600)
			if err != nil {
				log.Print(err)
	
			}
		}
		
}