//go:build !windows

package lang

/*
#include <sys/mman.h>
#include <unistd.h>

typedef struct {
	int32_t pkgPath;
	int16_t mcount;
	int16_t xcount;
	uint32_t Moff;
} UncommonType;

void patchExportMethods(void *ut) {
	long pageSize = sysconf(_SC_PAGESIZE);
	void *pageStart = (void *)((uintptr_t)ut & ~(pageSize - 1));
	mprotect(pageStart, pageSize, PROT_READ | PROT_WRITE);
	UncommonType *ptr = (UncommonType *)ut;
	ptr->xcount = ptr->mcount;
	mprotect(pageStart, pageSize, PROT_READ);
}
*/
import "C"
import "unsafe"

func patchExportMethods(ut unsafe.Pointer) {
	C.patchExportMethods(ut)
}
