package service

import (
	"errors"
	"tinnydouban/global"
	"tinnydouban/pkg/upload"
	"mime/multipart"
	"os"
)

//我们在 UploadFile Service 方法中，主要是通过获取文件所需的基本信息，
//接着对其进行业务所需的文件检查（文件大小是否符合需求、文件后缀是否达到要求），
//并且判断在写入文件前对否具备必要的写入条件（目录是否存在、权限是否足够），
//最后在检查通过后再进行真正的写入文件操作。

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported.")
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit.")
	}

	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}

	dst := uploadSavePath + "/" + fileName
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
