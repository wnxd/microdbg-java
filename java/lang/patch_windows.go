//go:build windows

package lang

/*
#include <windows.h>

typedef struct {
	INT32 pkgPath;
	UINT16 mcount;
	UINT16 xcount;
	UINT32 Moff;
} UncommonType;

void patchExportMethods(void *ut) {
	DWORD oldProtect;
	VirtualProtect(ut, sizeof(UncommonType), PAGE_READWRITE, &oldProtect);
	UncommonType *ptr = (UncommonType *)ut;
	ptr->xcount = ptr->mcount;
	VirtualProtect(ut, sizeof(UncommonType), oldProtect, NULL);
}
*/
import "C"
import "unsafe"

func patchExportMethods(ut unsafe.Pointer) {
	C.patchExportMethods(ut)
}
