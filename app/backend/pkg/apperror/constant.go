package apperror

import "net/http"

type Locale string

const (
	LocaleVi Locale = "vi"
	LocaleEn Locale = "en"
	LocaleKo Locale = "ko"
)

var (
	ErrInternalServerError = ErrorResponse{
		Code: http.StatusInternalServerError,
		Message: ErrorMessage{
			Vi: "Lỗi máy chủ nội bộ",
			En: "Internal server error",
			Ko: "내부 서버 오류",
		},
	}
	ErrCompanyIDRequired = ErrorResponse{
		Code: http.StatusBadRequest,
		Message: ErrorMessage{
			Vi: "Company ID là bắt buộc",
			En: "Company ID is required",
			Ko: "Company ID는 필수입니다",
		},
	}
	ErrUserIDRequired = ErrorResponse{
		Code: http.StatusBadRequest,
		Message: ErrorMessage{
			Vi: "User ID là bắt buộc",
			En: "User ID is required",
			Ko: "User ID는 필수입니다",
		},
	}
	ErrInvalidRequest = ErrorResponse{
		Code: http.StatusBadRequest,
		Message: ErrorMessage{
			Vi: "Yêu cầu không hợp lệ",
			En: "Invalid request",
			Ko: "잘못된 요청",
		},
	}
	ErrNotFound = ErrorResponse{
		Code: http.StatusNotFound,
		Message: ErrorMessage{
			Vi: "Không tìm thấy",
			En: "Not found",
			Ko: "찾을 수 없음",
		},
	}
	ErrUnauthorized = ErrorResponse{
		Code: http.StatusUnauthorized,
		Message: ErrorMessage{
			Vi: "Không có quyền truy cập",
			En: "Unauthorized",
			Ko: "권한이 없습니다",
		},
	}
	ErrForbidden = ErrorResponse{
		Code: http.StatusForbidden,
		Message: ErrorMessage{
			Vi: "Không có quyền truy cập",
			En: "Forbidden",
			Ko: "접근이 금지되었습니다",
		},
	}
	ErrLoginInformationInvalid = ErrorResponse{
		Code: http.StatusBadRequest,
		Message: ErrorMessage{
			Vi: "Thông tin đăng nhập không hợp lệ",
			En: "Login information is invalid",
			Ko: "로그인 정보가 유효하지 않습니다",
		},
	}
	ErrConfigAlreadyExists = ErrorResponse{
		Code: http.StatusBadRequest,
		Message: ErrorMessage{
			Vi: "Config đã tồn tại",
			En: "Config already exists",
			Ko: "Config가 이미 존재합니다",
		},
	}
)
