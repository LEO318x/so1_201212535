#include <linux/module.h>	
#include <linux/kernel.h>	
#include <linux/init.h>		
#include <linux/proc_fs.h>
#include <asm/uaccess.h>	
#include <linux/seq_file.h>
#include <linux/hugetlb.h>

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo de RAM SO1");
MODULE_AUTHOR("Mike Molina");

struct sysinfo inf;


static int write_file(struct seq_file *file , void *v)
{

	long memoriaTotal;
    long memoriaLibre; 
	long memoriaBuffer; 

	si_meminfo(&inf);
 	
	memoriaTotal = (inf.totalram * inf.mem_unit);
    memoriaLibre = (inf.freeram * inf.mem_unit); 
	memoriaBuffer = (inf.bufferram * inf.mem_unit); 
    seq_printf(file, "{\"MemTotal\":%li,\"MemLibre\":%li,\"MemBuffer\":%li}",memoriaTotal/1024/1024,memoriaLibre/1024/1024,memoriaBuffer/1024/1024);
	return 0;
}

static int to_open_file(struct inode *inode,struct file *file)
{
	return single_open(file,write_file,NULL);
}

static struct proc_ops operations = 
{
	.proc_open = to_open_file,
	.proc_read = seq_read
};

static int __init ram_201212535_init(void)
{
	proc_create("ram_201212535",0,NULL, &operations);
	printk(KERN_INFO "201212535 \n");
	return 0;
}

static void __exit ram_201212535_exit(void)
{
	remove_proc_entry("ram_201212535",NULL);
	printk(KERN_INFO "Sistemas Operativos 1\n");
}

module_init(ram_201212535_init);
module_exit(ram_201212535_exit);
